package router

import (
	"fmt"
	"net/http"

	"github.com/doanba13/job-view-service/models"
	"github.com/doanba13/job-view-service/storages"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllCatalog(context *fiber.Ctx) error {
	cats := &[]models.Catalog{}

	err := storages.Repo.DB.Find(cats).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success Catalog",
		"data":    cats,
	})

	return nil

}
func GetAllCatalogJob(context *fiber.Ctx) error {
	cats := &models.Catalog{}
	id := context.Params("id")

	fmt.Println(id)

	err := storages.Repo.DB.Model(&models.Catalog{}).Preload("Jobs").Where("id = ?", id).Find(cats).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success Catalog",
		"data":    cats,
	})

	return nil

}

func GetCatalogById(db *gorm.DB, id string) *models.Catalog {
	cats := &models.Catalog{}

	err := storages.Repo.DB.Find(cats, id).Error

	if err != nil {
		fmt.Printf("Cannot find catalog")
	}

	return cats
}
func GetCatalogByIdList(db *gorm.DB, catalogs []string) []models.Catalog {
	cats := &[]models.Catalog{}

	err := storages.Repo.DB.Where("name IN ?", catalogs).Find(cats).Error

	if err != nil {
		fmt.Printf("Cannot find catalog")
	}

	return *cats
}
