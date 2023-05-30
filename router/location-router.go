package router

import (
	"net/http"

	"github.com/doanba13/job-view-service/models"
	"github.com/doanba13/job-view-service/storages"
	"github.com/gofiber/fiber/v2"
)

func GetAllLocation(context *fiber.Ctx) error {
	locs := &[]models.Location{}

	err := storages.Repo.DB.Find(locs).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"data":    locs,
	})

	return nil

}
