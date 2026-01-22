package models

import "golang.org/x/crypto/bcrypt"

// Initialize data status in custom type DataStatus
type DataStatus string

const (
	Active   DataStatus = "active"
	Inactive DataStatus = "inactive"
)

type UserRole string

const (
	Operator      UserRole = "operator"
	Cashier       UserRole = "cashier"
	Finance       UserRole = "finance"
	Umum          UserRole = "umum"
	Superadmin    UserRole = "superadmin"
	Administrator UserRole = "administrator"
)

// Pendaftaran UserRole = "pendaftaran"
// Rekammedis  UserRole = "rekammedis"
// Ralan       UserRole = "ralan"
// Ranap       UserRole = "ranap"
// Vk          UserRole = "vk"
// Lab         UserRole = "lab"
// Klaim       UserRole = "klaim"
// Simrs       UserRole = "simrs"
// Ipsrs       UserRole = "ipsrs"

// Initialize user model
type User struct {
	ID         string     `gorm:"type:varchar(15);primaryKey" json:"id" validate:"required"`
	Username   string     `gorm:"type:varchar(255);not null;unique" json:"username" validate:"required"`
	Password   string     `gorm:"type:text;not null" json:"password" validate:"required"`
	Name       string     `gorm:"type:varchar(255);not null" json:"name" validate:"required"`
	UserRole   UserRole   `gorm:"type:user_role;not null;default:'operator'" json:"user_role" validate:"required"`
	UserStatus DataStatus `gorm:"type:data_status;not null;default:'inactive'" json:"user_status" validate:"required"`
}

// SetID is function to set ID into User
func (b *User) SetID(id string) {
	b.ID = id
}

// HashPassword is a function to hash password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
