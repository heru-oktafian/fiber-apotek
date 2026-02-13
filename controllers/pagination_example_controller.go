package controllers

import (
	fiber "github.com/gofiber/fiber/v2"
	configs "github.com/heru-oktafian/fiber-apotek/configs"
	helpers "github.com/heru-oktafian/fiber-apotek/helpers"
	models "github.com/heru-oktafian/fiber-apotek/models"
	services "github.com/heru-oktafian/fiber-apotek/services"
)

// ==================== CONTOH IMPLEMENTASI DENGAN PaginateWithSearchAndMonth ====================

// GetAllBuyReturnsSimplified adalah contoh refactored dari GetAllBuyReturns menggunakan helper
func GetAllBuyReturnsSimplified(c *fiber.Ctx) error {
	branchID, _ := services.GetBranchID(c)

	var buyReturnsFromDB []models.AllBuyReturns

	// Konfigurasi query dasar
	query := configs.DB.Table("buy_returns A").
		Select("A.id, A.purchase_id, A.return_date, A.payment, A.total_return").
		Where("A.branch_id = ? ", branchID).
		Order("A.created_at DESC")

	// Gunakan helper untuk pagination, search, dan month
	data, search, total, page, totalPages, err := helpers.PaginateWithSearchAndMonth(
		c,
		query,
		&buyReturnsFromDB,
		[]string{"A.purchase_id"}, // searchColumn: kolom mana yang akan di-search
		"A.return_date",           // dateColumn: kolom tanggal untuk filter bulan
		1,                         // defaultPage: halaman awal jika tidak disediakan
		10,                        // defaultLimit: 10 data per halaman
	)

	if err != nil {
		return helpers.JSONResponse(c, fiber.StatusBadRequest, "Parameter tidak valid", err.Error())
	}

	// Cast data ke slice models.AllBuyReturns
	formattedData := data.([]models.AllBuyReturns)

	// Format data sebelum mengirim response
	var formattedBuyReturnsData []models.BuyReturnsResponse
	for _, buyReturn := range formattedData {
		formattedBuyReturnsData = append(formattedBuyReturnsData, models.BuyReturnsResponse{
			ID:          buyReturn.ID,
			PurchaseId:  buyReturn.PurchaseId,
			ReturnDate:  helpers.FormatIndonesianDate(buyReturn.ReturnDate),
			TotalReturn: buyReturn.TotalReturn,
			Payment:     string(buyReturn.Payment),
		})
	}

	return helpers.JSONResponseGetAll(
		c,
		fiber.StatusOK,
		"Data retur pembelian berhasil diambil",
		search,
		total,
		page,
		totalPages,
		10,
		formattedBuyReturnsData,
	)
}

// ==================== CONTOH DENGAN MULTIPLE SEARCH COLUMNS ====================

// GetAllBuyReturnsWithMultiSearch contoh jika ingin search di multiple columns
// NOTE: Helper saat ini hanya support satu column, untuk multiple columns bisa di-handle manual sebelum helper
func GetAllBuyReturnsWithMultiSearch(c *fiber.Ctx) error {
	branchID, _ := services.GetBranchID(c)
	search := c.Query("search")

	var buyReturnsFromDB []models.AllBuyReturns

	query := configs.DB.Table("buy_returns A").
		Select("A.id, A.purchase_id, A.return_date, A.payment, A.total_return").
		Where("A.branch_id = ? ", branchID).
		Order("A.created_at DESC")

	// Gunakan helper untuk pagination, search, dan month (support multi-column search)
	data, _, total, page, totalPages, err := helpers.PaginateWithSearchAndMonth(
		c,
		query,
		&buyReturnsFromDB,
		[]string{"A.purchase_id", "A.payment"}, // Search di kedua kolom ini
		"A.return_date",
		1,
		10,
	)

	if err != nil {
		return helpers.JSONResponse(c, fiber.StatusBadRequest, "Parameter tidak valid", err.Error())
	}

	formattedData := data.([]models.AllBuyReturns)

	var formattedBuyReturnsData []models.BuyReturnsResponse
	for _, buyReturn := range formattedData {
		formattedBuyReturnsData = append(formattedBuyReturnsData, models.BuyReturnsResponse{
			ID:          buyReturn.ID,
			PurchaseId:  buyReturn.PurchaseId,
			ReturnDate:  helpers.FormatIndonesianDate(buyReturn.ReturnDate),
			TotalReturn: buyReturn.TotalReturn,
			Payment:     string(buyReturn.Payment),
		})
	}

	return helpers.JSONResponseGetAll(
		c,
		fiber.StatusOK,
		"Data retur pembelian berhasil diambil",
		search,
		total,
		page,
		totalPages,
		10,
		formattedBuyReturnsData,
	)
}

// ==================== CONTOH DENGAN LIMIT BERBEDA ====================

// GetAllBuyReturnsWithCustomLimit contoh menggunakan limit 20 per halaman
func GetAllBuyReturnsWithCustomLimit(c *fiber.Ctx) error {
	branchID, _ := services.GetBranchID(c)

	var buyReturnsFromDB []models.AllBuyReturns

	query := configs.DB.Table("buy_returns A").
		Select("A.id, A.purchase_id, A.return_date, A.payment, A.total_return").
		Where("A.branch_id = ? ", branchID).
		Order("A.created_at DESC")

	// Helper dengan limit 20
	data, search, total, page, totalPages, err := helpers.PaginateWithSearchAndMonth(
		c,
		query,
		&buyReturnsFromDB,
		[]string{"A.purchase_id"},
		"A.return_date",
		1,
		20, // Limit 20 per halaman
	)

	if err != nil {
		return helpers.JSONResponse(c, fiber.StatusBadRequest, "Parameter tidak valid", err.Error())
	}

	formattedData := data.([]models.AllBuyReturns)

	var formattedBuyReturnsData []models.BuyReturnsResponse
	for _, buyReturn := range formattedData {
		formattedBuyReturnsData = append(formattedBuyReturnsData, models.BuyReturnsResponse{
			ID:          buyReturn.ID,
			PurchaseId:  buyReturn.PurchaseId,
			ReturnDate:  helpers.FormatIndonesianDate(buyReturn.ReturnDate),
			TotalReturn: buyReturn.TotalReturn,
			Payment:     string(buyReturn.Payment),
		})
	}

	return helpers.JSONResponseGetAll(
		c,
		fiber.StatusOK,
		"Data retur pembelian berhasil diambil",
		search,
		total,
		page,
		totalPages,
		20, // Sesuaikan dengan limit yang digunakan
		formattedBuyReturnsData,
	)
}

// ==================== CONTOH UNTUK PURCHASE DATA ====================

// GetAllPurchasesWithPagination contoh implementasi untuk tabel purchases
func GetAllPurchasesWithPagination(c *fiber.Ctx) error {
	branchID, _ := services.GetBranchID(c)

	var purchases []struct {
		ID            string `json:"id"`
		PurchaseDate  string `json:"purchase_date"`
		SupplierName  string `json:"supplier_name"`
		TotalPurchase int    `json:"total_purchase"`
	}

	query := configs.DB.Table("purchases A").
		Select("A.id, A.purchase_date, B.name AS supplier_name, A.total_purchase").
		Joins("LEFT JOIN suppliers B ON B.id = A.supplier_id").
		Where("A.branch_id = ?", branchID).
		Order("A.purchase_date DESC")

	// Gunakan helper untuk pagination, search, dan month
	data, search, total, page, totalPages, err := helpers.PaginateWithSearchAndMonth(
		c,
		query,
		&purchases,
		[]string{"A.id"},  // Search berdasarkan purchase ID
		"A.purchase_date", // Filter berdasarkan tanggal pembelian
		1,
		10,
	)

	if err != nil {
		return helpers.JSONResponse(c, fiber.StatusBadRequest, "Parameter tidak valid", err.Error())
	}

	return helpers.JSONResponseGetAll(
		c,
		fiber.StatusOK,
		"Data pembelian berhasil diambil",
		search,
		total,
		page,
		totalPages,
		10,
		data,
	)
}

// ==================== TIPS DEBUGGING ====================

// TipsDebugging - Pastikan mengikuti langkah-langkah berikut:
//
// 1. Cek apakah model sudah di-initialize :
//    ✓ var buyReturnsFromDB []models.AllBuyReturns
//    ✗ var buyReturnsFromDB []models.AllBuyReturns = nil (tidak perlu, otomatis jadi slice kosong)
//
// 2. Gunakan pointer saat memanggil helper:
//    ✓ helpers.PaginateWithSearchAndMonth(c, query, &buyReturnsFromDB, ...)
//    ✗ helpers.PaginateWithSearchAndMonth(c, query, buyReturnsFromDB, ...)
//
// 3. Cast data dengan benar:
//    ✓ formattedData := data.([]models.AllBuyReturns)
//    ✗ formattedData := data (tidak akan bisa iterate)
//
// 4. Query harus sudah include Select dan Where:
//    ✓ query := configs.DB.Table("...").Select("...").Where("...")
//    ✗ query := configs.DB.Table("...") (akan query semua kolom)
//
// 5. Pastikan searchColumn dan dateColumn sesuai dengan konfigurasi query:
//    ✓ query.Select("A.id, A.purchase_id, A.return_date") + PaginateWithSearchAndMonth(..., "A.purchase_id", "A.return_date", ...)
//    ✗ query.Select("A.id") + PaginateWithSearchAndMonth(..., "A.purchase_id", "A.return_date", ...) (kolom tidak ada)
