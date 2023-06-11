package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/doanba13/job-view-service/models"
	"github.com/doanba13/job-view-service/storages"
	"github.com/gofiber/fiber/v2"
)

type Job struct {
	Id          string `json:"id"`
	CompanyId   string `json:"companyId"`
	UserId      string `json:"userId"`
	Title       string `json:"title"`
	Introduce   string `json:"introduce"`
	Desc        string `json:"desc"`
	Exp         string `json:"exp"`
	Salary      int    `json:"salary"`
	DateCreated int    `json:"dateCreated"`
	Unit        int    `json:"unit"`
	View        int    `json:"view"`
	Catalog     []int  `json:"catalog"`
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

func SearchJob(context *fiber.Ctx) error {
	jobModel := &[]models.JobEntity{}
	search := strings.ToLower(context.Query("keyword"))

	fmt.Println(search)

	err := storages.Repo.DB.Preload("Catalogs").Where("LOWER(title) like ?", "%"+search+"%").Or("LOWER(introduce) like ?", "%"+search+"%").Or("LOWER(exp) like ?", "%"+search+"%").Find(jobModel).Error

	if err != nil {
		fmt.Println(err.Error())
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"data":    jobModel,
	})

	return nil
}

func GetJobByView(context *fiber.Ctx) error {
	jobModel := &[]models.JobEntity{}

	err := storages.Repo.DB.Limit(10).Preload("Catalogs").Order("view DESC").Find(jobModel).Error

	if err != nil {
		fmt.Println(err.Error())
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"data":    jobModel,
	})

	return nil
}

func GetHighSalaryJob(context *fiber.Ctx) error {
	jobModel := &[]models.JobEntity{}

	err := storages.Repo.DB.Limit(10).Preload("Catalogs").Order("salary DESC").Find(jobModel).Error

	if err != nil {
		fmt.Println(err.Error())
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success",
		"data":    jobModel,
	})
	return nil
}

func GetRecentlyJob(context *fiber.Ctx) error {
	jobModel := &[]models.JobEntity{}

	err := storages.Repo.DB.Limit(10).Preload("Catalogs").Order("date_created DESC").Find(jobModel).Error

	if err != nil {
		fmt.Println(err.Error())
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

	err := storages.Repo.DB.Where("id = ?", id).Preload("Catalogs").Find(jobModel).Error

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
