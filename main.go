package main

import (
	"github.com/Thivanka/ABC_Pharmacy/server/bootstrap"
	"github.com/Thivanka/ABC_Pharmacy/server/repository"
	"github.com/gofiber/fiber/v2"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	bootstrap.InitializeApp(app)

}
