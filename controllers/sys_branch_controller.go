package controllers

import (
	fiber "github.com/gofiber/fiber/v2"
	config "github.com/heru-oktafian/fiber-apotek/configs"
	helpers "github.com/heru-oktafian/fiber-apotek/helpers"
	models "github.com/heru-oktafian/fiber-apotek/models"
	services "github.com/heru-oktafian/fiber-apotek/services"
)

// CreateBranch is function for create new branch
func CreateBranch(c *fiber.Ctx) error {
	// Creating new unit using helpers
	return helpers.CreateResource(c, config.DB, &models.Branch{}, "BRC")
}

// UpdateBranch is function for update branch
func UpdateBranch(c *fiber.Ctx) error {
	id := c.Params("id")
	// Updating branch using helpers
	return helpers.UpdateResource(c, config.DB, &models.Branch{}, id)
}

// DeleteBranch is function for delete branch
func DeleteBranch(c *fiber.Ctx) error {
	id := c.Params("id")
	// Deleting branch using helpers
	return helpers.DeleteResource(c, config.DB, &models.Branch{}, id)
}

// GetBranch is function for get branch
func GetBranch(c *fiber.Ctx) error {
	id := c.Params("id")
	// Getting branch using helpers
	return helpers.GetResource(c, config.DB, &models.Branch{}, id)
}

// GetAllBranch is function for get all branch
func GetAllBranch(c *fiber.Ctx) error {
	var branches []models.Branch
	// Getting all branches using helpers
	return helpers.GetAllBranches(c, config.DB, &branches)
}

// GetUserBranch menangani penampilan userbranch
func GetBranchByUserId(c *fiber.Ctx) error {
	// get user id
	userID, _ := services.GetUserID(c)

	// Menampilkan semua userbranch
	var userBranchDetails []models.UserBranchDetail

	// Melakukan LEFT OUTER JOIN menggunakan GORM
	if err := config.DB.
		Table("user_branches").
		Select("user_branches.user_id, users.name AS user_name, user_branches.branch_id, branches.branch_name, branches.sia_name, branches.sipa_name, branches.phone").
		Joins("LEFT JOIN users ON users.id = user_branches.user_id").
		Joins("LEFT JOIN branches ON branches.id = user_branches.branch_id").
		Where("branches.branch_status = 'active' AND user_branches.user_id = ?", userID).
		Scan(&userBranchDetails).Error; err != nil {
		return helpers.JSONResponse(c, fiber.StatusInternalServerError, "Get userbranches failed", "Failed to fetch user branches with details")
	}

	// Mengembalikan response data userbranch
	return helpers.JSONResponse(c, fiber.StatusOK, "User Branch found", userBranchDetails)
}
