package repository

import (
	"net/http"

	"github.com/Thivanka/ABC_Pharmacy/server/database/migrations"
	"github.com/Thivanka/ABC_Pharmacy/server/database/models"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gopkg.in/go-playground/validator.v9"
)

type ErrorResponseInvoice struct {
	FaildField string
	Tag        string
	value      string
}

var validateinvoice = validator.New()

func ValidateStructinvoice(invoice models.Invoice) []*ErrorResponseInvoice {
	var errors []*ErrorResponseInvoice
	err := validateinvoice.Struct(invoice)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponseInvoice
			element.FaildField = err.StructNamespace()
			element.Tag = err.Tag()
			element.value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (r *Repository) GetInvoice(context *fiber.Ctx) error {
	db := r.DB
	model := db.Model(&migrations.Invoices{})

	pg := paginate.New(&paginate.Config{
		DefaultSize:        20,
		CustomParamEnabled: true,
	})

	page := pg.With(model).Request(context.Request()).Response(&[]migrations.Invoices{})

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"data": page,
	})

	return nil
}

func (r *Repository) CreateInvoice(context *fiber.Ctx) error {
	invoice := models.Invoice{}
	err := context.BodyParser(&invoice)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request faild"})

		return err
	}

	errors := ValidateStructinvoice(invoice)

	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if err := r.DB.Create(invoice).Error; err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create invoice",
			"data":    err,
		})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Invoice has been added!",
		"data":    invoice,
	})

	return nil
}

func (r *Repository) UpdateInvoice(context *fiber.Ctx) error {
	invoice := models.Invoice{}
	err := context.BodyParser(&invoice)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request faild"})

		return err
	}

	errors := ValidateStructinvoice(invoice)
	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db := r.DB
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id can not be empty !",
		})
		return nil
	}
	if db.Model(&invoice).Where("id = ?", id).Updates(&invoice).RowsAffected == 0 {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not get invoice profile !",
		})
		return nil
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Invoice sucessfully updated !",
	})
	return nil
}

func (r *Repository) DeleteInvoice(context *fiber.Ctx) error {
	invoiceModel := &migrations.Invoices{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id can not be empty !",
		})
		return nil
	}

	err := r.DB.Delete(invoiceModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not delete !",
		})

		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Invoices sucessfully Deleted !",
	})
	return nil
}

func (r *Repository) GetInvoiceById(context *fiber.Ctx) error {
	invoiceModel := &migrations.Invoices{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id can not be empty !",
		})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(invoiceModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not the get the invoice !",
		})

		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Invoice profile fetch sucessfully !",
		"data":    invoiceModel,
	})
	return nil
}
