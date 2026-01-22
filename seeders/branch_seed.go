package seeders

import (
	time "time"

	config "github.com/heru-oktafian/fiber-apotek/configs"
	models "github.com/heru-oktafian/fiber-apotek/models"
)

func BranchSeed() {
	t := time.Date(2026, time.December, 31, 0, 0, 0, 0, time.UTC)

	branch := []models.Branch{
		{ID: "BRC0551960Y9TY0", BranchName: "Abi Foundation", Address: "Jl. Raya Gudo, No. 101A, Kecamatan Gudo, Kabupaten Jombang, Jawa Timur", Phone: "085236990001", Email: "info@heruoktafian.com", OwnerId: "3517011710880001", OwnerName: "Heru Oktafian, ST., CTT", BankName: "Bank BCA", AccountName: "Heru Oktafian", AccountNumber: "1520582106", JournalMethod: "automatic", TaxPercentage: 11, BranchStatus: "active", LicenseDate: t},
	}
	config.DB.Create(&branch)
}
