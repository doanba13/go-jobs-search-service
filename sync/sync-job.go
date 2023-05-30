package sync

import (
	"encoding/json"
	"fmt"

	"github.com/doanba13/job-view-service/models"
	"github.com/doanba13/job-view-service/router"

	"gorm.io/gorm"
)

// type Job struct {
// 	Id        string           `json:"_id"`
// 	CompanyId string           `json:"companyId"`
// 	UserId    string           `json:"userId"`
// 	Title     string           `json:"title"`
// 	Introduce string           `json:"introduce"`
// 	Desc      string           `json:"desc"`
// 	Exp       string           `json:"exp"`
// 	Salary    string           `json:"salary"`
// 	Unit      int              `json:"unit"`
// 	View      int              `json:"view"`
// 	Catalog   []models.Catalog `gorm:"many2many:job_catalog;foreignKey:Id" json:"catalog"`
// }

type JobData struct {
	Info       models.JobEntity `json:"info"`
	CatalogsId []string         `json:"catalogs"`
}

func SaveJob(db *gorm.DB, data string) {
	c := JobData{}
	err := json.Unmarshal([]byte(data), &c)
	fmt.Println("========================================")
	fmt.Println(c.CatalogsId)

	cats := router.GetCatalogByIdList(db, c.CatalogsId)

	if err != nil {
		fmt.Println(err)
	}

	c.Info.Catalogs = cats
	fmt.Println("========================================")
	fmt.Println(&c.Info)
	fmt.Println("========================================")

	err = db.Create(&c.Info).Error

	if err != nil {
		fmt.Println(err)
	}
}
