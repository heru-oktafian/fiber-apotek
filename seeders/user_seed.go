package seeders

import (
	config "github.com/heru-oktafian/fiber-apotek/configs"
	models "github.com/heru-oktafian/fiber-apotek/models"
)

func UserSeed() {
	users := []models.User{
		{ID: "USR250118132201", Username: "superadmin", Password: "Superadmin123", Name: "Super Admin", UserRole: "superadmin", UserStatus: "active"},
		{ID: "USR250118132202", Username: "operator", Password: "Operator123", Name: "Operator", UserRole: "operator", UserStatus: "active"},
	}

	// Hash password for each user
	for _, user := range users {
		if err := user.HashPassword(); err != nil {
			continue
		}
		config.DB.Create(&user)
	}
}
