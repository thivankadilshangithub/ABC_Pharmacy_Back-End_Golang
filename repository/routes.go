package repository

import (
	"github.com/gofiber/fiber/v2"
)

// Items routers
func (repo *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Items routers
	api.Get("/items", repo.GetItems)
	api.Post("/items", repo.CreateItems)
	api.Patch("/items/:id", repo.UpdateItem)
	api.Delete("/items/:id", repo.DeleteItem)
	api.Get("/items/:id", repo.GetItemById)

	//Invoices routers
	api.Get("/invoices", repo.GetInvoice)
	api.Post("/invoices", repo.CreateInvoice)
	api.Patch("/invoices/:id", repo.UpdateInvoice)
	api.Delete("/invoices/:id", repo.DeleteInvoice)
	api.Get("/invoices/:id", repo.GetInvoiceById)
}
