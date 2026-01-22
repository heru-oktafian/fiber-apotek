package services

import (
	fiber "github.com/gofiber/fiber/v2"
	helpers "github.com/heru-oktafian/fiber-apotek/helpers"
)

func GetUserID(c *fiber.Ctx) (string, error) {
	// Get user_id from claims
	userID, _ := helpers.GetClaimsToken(c, "sub")
	// println("+" + userID)
	return userID, nil
}

func GetBranchID(c *fiber.Ctx) (string, error) {
	// Get branch_id from claims
	branchID, _ := helpers.GetClaimsToken(c, "branch_id")

	return branchID, nil
}

func GetUserRole(c *fiber.Ctx) (string, error) {
	// Get user_role from claims
	userRole, _ := helpers.GetClaimsToken(c, "user_role")

	return userRole, nil
}
