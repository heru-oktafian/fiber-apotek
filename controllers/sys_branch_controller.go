package controllers

import (
	fiber "github.com/gofiber/fiber/v2"
	config "github.com/heru-oktafian/fiber-apotek/configs"
	helpers "github.com/heru-oktafian/fiber-apotek/helpers"
	models "github.com/heru-oktafian/fiber-apotek/models"
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
