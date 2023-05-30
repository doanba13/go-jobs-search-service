package router

import (
	"log"
	"net/http"

	"github.com/doanba13/job-view-service/models"
	"github.com/doanba13/job-view-service/storages"
	"github.com/gofiber/fiber/v2"
)

type Job struct {
	Id        string `json:"id"`
	CompanyId string `json:"companyId"`
	UserId    string `json:"userId"`
	Title     string `json:"title"`
	Introduce string `json:"introduce"`
	Desc      string `json:"desc"`
	Exp       string `json:"exp"`
	Salary    string `json:"salary"`
	Unit      int    `json:"unit"`
	View      int    `json:"view"`
	Catalog   []int  `json:"catalog"`
}

func CreateJob(context *fiber.Ctx) error {
	job := Job{}

	err := context.BodyParser(&job)

	if err != nil {
		log.Fatal("Cannot decode json")
		return err
	}

	err = storages.Repo.DB.Create(job).Error

	if err != nil {
		log.Fatal("Cannot create job")
		return err
	}
	return nil
}

func GetAllJob(context *fiber.Ctx) error {
	jobModel := &[]models.JobEntity{}

	err := storages.Repo.DB.Preload("Catalogs").Find(jobModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"data":    jobModel,
	})

	return nil

}

func GetAllJobByCompanyId(context *fiber.Ctx) error {
	jobModel := &[]models.JobEntity{}
	companyId := context.Query("companyId")

	err := storages.Repo.DB.Where("companyId = ?", companyId).Preload("Catalogs").Find(jobModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"data":    jobModel,
	})

	return nil

}

func DetailJob(context *fiber.Ctx) error {
	jobModel := &models.JobEntity{}
	id := context.Params("id")

	err := storages.Repo.DB.Find(jobModel, id).Error

	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "Job not found!"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"data":    jobModel,
	})

	return nil

}
