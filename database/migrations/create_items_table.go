package migrations

import (
	"gorm.io/gorm"
)

type Items struct {
	ID           uint    `gorm:"primary key:autoIncrement" json: "id" `
	Name         *string `json:"name"`
	UnitPrice    *string `json:"unitprice"`
	ItemCategory *string `json:"itemcategory"`
}

func MigrateItems(db *gorm.DB) error {
	err := db.AutoMigrate(&Items{})
	return err
}
