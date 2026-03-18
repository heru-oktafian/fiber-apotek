package controllers

import (
	fmt "fmt"
	strings "strings"

	fiber "github.com/gofiber/fiber/v2"
	configs "github.com/heru-oktafian/fiber-apotek/configs"
	helpers "github.com/heru-oktafian/fiber-apotek/helpers"
	models "github.com/heru-oktafian/fiber-apotek/models"
	services "github.com/heru-oktafian/fiber-apotek/services"
)

// CreateProduct buat Product
func CreateProduct(c *fiber.Ctx) error {
	// Creating new Product using helpers
	return helpers.CreateResource(c, configs.DB, &models.Product{}, "PRD")
}

// UpdateProduct update Product
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	// Updating Product using helpers
	return helpers.UpdateResource(c, configs.DB, &models.Product{}, id)
}

// DeleteProduct hapus Product
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	// Deleting Product using helpers
	return helpers.DeleteResource(c, configs.DB, &models.Product{}, id)
}

// GetProduct tampilkan Product berdasarkan id
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var AllProduct []models.ProductDetail
	if err := configs.DB.
		Table("products pro").
		Select("pro.id,pro.sku,pro.name,pro.description, pro.ingredient, pro.dosage, pro.side_affection, pro.unit_id AS unit_id,pro.stock,pro.purchase_price,pro.expired_date,pro.sales_price, pro.alternate_price, pro.product_category_id,pc.name AS product_category_name,un.name AS unit_name,pro.branch_id").
		Joins("LEFT JOIN product_categories pc ON pc.id = pro.product_category_id").
		Joins("LEFT JOIN units un ON un.id = pro.unit_id").
		Where("pro.id = ?", id).
		Scan(&AllProduct).Error; err != nil {
		return helpers.JSONResponse(c, fiber.StatusNotFound, "Data tidak ditemukan", err)
	}

	// print(AllProduct)
	return helpers.JSONResponse(c, fiber.StatusOK, "Data ditemukan", AllProduct)
}

// GetAllProduct tampilkan semua Product
func GetAllProduct(c *fiber.Ctx) error {
	// Ambil ID cabang
	branch_id, _ := services.GetBranchID(c)

	var AllProduct []models.ProductDetail
	var total int

	// Query dasar
	query := configs.DB.Table("products pro").
		Select("pro.id,pro.sku,pro.name,pro.description, pro.ingredient, pro.dosage, pro.side_affection, pro.unit_id, un.name AS unit_name,pro.stock,pro.purchase_price,pro.sales_price,pro.alternate_price,pro.expired_date, pro.product_category_id, pc.name AS product_category_name").
		Joins("LEFT JOIN product_categories pc ON pc.id = pro.product_category_id").
		Joins("LEFT JOIN units un ON un.id = pro.unit_id").
		Where("pro.branch_id = ?", branch_id)

	_, search, total, page, totalPages, limit, err := helpers.Paginate(c, query, &AllProduct, []string{"pro.name, pro.alias, pro.description, pro.ingredient, pro.dosage, pro.side_affection"})
	if err != nil {
		return helpers.JSONResponse(c, fiber.StatusInternalServerError, "Get AllProduct failed", err.Error())
	}

	return helpers.JSONResponseGetAll(c, fiber.StatusOK, "Products retrieved successfully", search, int(total), page, int(totalPages), int(limit), AllProduct)

}

// CmbProdSale mengembalikan daftar produk untuk combo box transaksi penjualan
func CmbProdSale(c *fiber.Ctx) error {
	branch_id, _ := services.GetBranchID(c)
	user_id, _ := services.GetUserID(c)
	search := strings.TrimSpace(c.Query("search"))

	// Buat cacheKey berdasarkan branch_id dan user_id
	cacheKey := fmt.Sprintf("%v:%v", branch_id, user_id)

	// Cek apakah ada data di Redis terlebih dahulu
	cachedProducts, err := services.GetTemporaryProductCache(cacheKey)
	if err != nil {
		fmt.Printf("Failed to get product cache for cacheKey %s: %v\n", cacheKey, err)
		// Lanjutkan ke query database jika gagal ambil cache
	}
	if cachedProducts != nil {
		// Jika ada data di cache, gunakan data tersebut
		return helpers.JSONResponse(c, fiber.StatusOK, "Combo Products retrieved from cache successfully", cachedProducts)
	}

	// Jika tidak ada di cache, lakukan query ke database
	var cmbProducts []models.ProdSaleCombo

	query := configs.DB.Table("products").
		Select("products.id as product_id, products.name as product_name, sales_price AS price, products.stock, products.unit_id, units.name AS unit_name").
		Joins("LEFT JOIN units ON units.id = products.unit_id").
		Where("products.branch_id = ?", branch_id)

	//if search != "" {
		search = strings.ToLower(search)
		query = query.Where("LOWER(products.name) LIKE ? OR LOWER(products.description) LIKE ? OR LOWER(products.id) LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	//}

	query = query.Order("products.name ASC")

	if err := query.Scan(&cmbProducts).Error; err != nil {
		return helpers.JSONResponse(c, fiber.StatusInternalServerError, "Get Combo Products failed", err)
	}

	// Simpan list produk ke Redis dengan cacheKey
	if err := services.SetTemporaryProductCache(cacheKey, cmbProducts); err != nil {
		fmt.Printf("Failed to save product cache for cacheKey %s: %v\n", cacheKey, err)
	}

	return helpers.JSONResponse(c, fiber.StatusOK, "Combo Products retrieved successfully", cmbProducts)
}

// CmbProdPurchase mengembalikan daftar produk untuk combo box transaksi pembelian
func CmbProdPurchase(c *fiber.Ctx) error {
	branch_id, _ := services.GetBranchID(c)
	user_id, _ := services.GetUserID(c)
	search := strings.TrimSpace(c.Query("search"))

	// Buat cacheKey berdasarkan branch_id dan user_id
	cacheKey := fmt.Sprintf("%v:%v", branch_id, user_id)

	// Cek apakah ada data di Redis terlebih dahulu
	cachedProducts, err := services.GetTemporaryPurchaseProductCache(cacheKey)
	if err != nil {
		fmt.Printf("Failed to get purchase product cache for cacheKey %s: %v\n", cacheKey, err)
		// Lanjutkan ke query database jika gagal ambil cache
	}
	if cachedProducts != nil {
		// Jika ada data di cache, gunakan data tersebut
		return helpers.JSONResponse(c, fiber.StatusOK, "Combo Purchase Products retrieved from cache successfully", cachedProducts)
	}

	// Jika tidak ada di cache, lakukan query ke database
	var cmbProducts []models.ProdPurchaseCombo

	query := configs.DB.Table("products").
		Select("products.id as product_id, products.name as product_name, purchase_price AS price, products.unit_id, units.name AS unit_name").
		Joins("LEFT JOIN units ON units.id = products.unit_id").
		Where("products.branch_id = ?", branch_id)

	if search != "" {
		search = strings.ToLower(search)
		query = query.Where("LOWER(products.name) LIKE ? OR LOWER(products.description) LIKE ? OR LOWER(products.id) LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	query = query.Order("products.name ASC")

	if err := query.Scan(&cmbProducts).Error; err != nil {
		return helpers.JSONResponse(c, fiber.StatusInternalServerError, "Get Combo Purchase Products failed", err)
	}

	// Simpan list produk ke Redis dengan cacheKey
	if err := services.SetTemporaryPurchaseProductCache(cacheKey, cmbProducts); err != nil {
		fmt.Printf("Failed to save purchase product cache for cacheKey %s: %v\n", cacheKey, err)
	}

	return helpers.JSONResponse(c, fiber.StatusOK, "Combo Purchase Products retrieved successfully", cmbProducts)
}
