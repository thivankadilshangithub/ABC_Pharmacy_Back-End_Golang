package migrations

import (
	"gorm.io/gorm"
)

type Invoices struct {
	ID         uint    `gorm:"primary key:autoIncrement" json: "id" `
	Name       *string `json:"name"`
	MobileNo   *string `json:"mobileno"`
	Email      *string `json:"email"`
	Address    *string `json:"address"`
	BilingType *string `json:"bilingtype"`
}

func MigrateInvoices(db *gorm.DB) error {
	err := db.AutoMigrate(&Invoices{})
	return err
}
