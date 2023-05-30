package router

import (
	"net/http"

	"github.com/doanba13/job-view-service/models"
	"github.com/doanba13/job-view-service/storages"
	"github.com/gofiber/fiber/v2"
)

func GetAllCompany(context *fiber.Ctx) error {
	comp := &[]models.Company{}

	err := storages.Repo.DB.Find(comp).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"data":    comp,
	})

	return nil

}
