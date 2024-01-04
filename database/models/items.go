package models

type Item struct {
	Name         string `json:"name" validate:"required,min=3,max=40" `
	UnitPrice    string `json:"unitprice" validate:"required,min=1,max=40" `
	ItemCategory string `json:"itemcategory" validate:"required,min=1,max=40" `
}
