package seeders

import (
	config "github.com/heru-oktafian/fiber-apotek/configs"
	models "github.com/heru-oktafian/fiber-apotek/models"
)

func UserBranchSeed() {
	userBranch := []models.UserBranch{
		{UserID: "USR250118132201", BranchID: "BRC0551960Y9TY0"},
		{UserID: "USR250118132202", BranchID: "BRC0551960Y9TY0"},
	}
	config.DB.Create(&userBranch)
}
