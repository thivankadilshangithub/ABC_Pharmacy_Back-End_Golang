package repository

import (
	"net/http"

	"github.com/Thivanka/ABC_Pharmacy/server/database/migrations"
	"github.com/Thivanka/ABC_Pharmacy/server/database/models"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gopkg.in/go-playground/validator.v9"
)

type ErrorResponse struct {
	FaildField string
	Tag        string
	value      string
}

var validate = validator.New()

func ValidateStruct(item models.Item) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(item)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FaildField = err.StructNamespace()
			element.Tag = err.Tag()
			element.value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (r *Repository) GetItems(context *fiber.Ctx) error {
	db := r.DB
	model := db.Model(&migrations.Items{})

	pg := paginate.New(&paginate.Config{
		DefaultSize:        20,
		CustomParamEnabled: true,
	})

	page := pg.With(model).Request(context.Request()).Response(&[]migrations.Items{})

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"data": page,
	})

	return nil
}

func (r *Repository) CreateItems(context *fiber.Ctx) error {
	item := models.Item{}
	err := context.BodyParser(&item)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request faild"})

		return err
	}

	errors := ValidateStruct(item)

	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if err := r.DB.Create(item).Error; err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create item",
			"data":    err,
		})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Item has been added!",
		"data":    item,
	})

	return nil
}

func (r *Repository) UpdateItem(context *fiber.Ctx) error {
	item := models.Item{}
	err := context.BodyParser(&item)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request faild"})

		return err
	}

	errors := ValidateStruct(item)
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
	if db.Model(&item).Where("id = ?", id).Updates(&item).RowsAffected == 0 {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not get item profile !",
		})
		return nil
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Item sucessfully updated !",
	})
	return nil
}

func (r *Repository) DeleteItem(context *fiber.Ctx) error {
	itemModel := &migrations.Items{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id can not be empty !",
		})
		return nil
	}

	err := r.DB.Delete(itemModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not delete !",
		})

		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Item sucessfully Deleted !",
	})
	return nil
}

func (r *Repository) GetItemById(context *fiber.Ctx) error {
	itemModel := &migrations.Items{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Id can not be empty !",
		})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(itemModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not the get the item !",
		})

		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Item profile fetch sucessfully !",
		"data":    itemModel,
	})
	return nil
}
