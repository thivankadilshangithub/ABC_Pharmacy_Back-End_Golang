package models

type Invoice struct {
	Name       string `json:"name" validate:"required,min=3,max=40" `
	MobileNo   string `json:"mobileno" validate:"required,min=5,max=40"`
	Email      string `json:"email" validate:"required,email,min=6,max=32"`
	Address    string `json:"address" validate:"required,min=5,max=60" `
	BilingType string `json:"bilingtype" validate:"required,min=3,max=40" `
}
