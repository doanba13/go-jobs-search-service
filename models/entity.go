package models

import "gorm.io/gorm"

type Catalog struct {
	ID   string      `gorm:"primary key" json:"id"`
	Name string      `json:"name"`
	Jobs []JobEntity `gorm:"many2many:job_catalog;foreignKey:ID" json:"jobs"`
}

type JobEntity struct {
	Id        string    `gorm:"primary key" json:"id"`
	CompanyId string    `json:"companyId"`
	UserId    string    `json:"userId"`
	Title     string    `json:"title"`
	Introduce string    `json:"introduce"`
	Desc      string    `json:"desc"`
	Exp       string    `json:"exp"`
	Salary    string    `json:"salary"`
	Unit      int       `json:"unit"`
	View      int       `json:"view"`
	Catalogs  []Catalog `gorm:"many2many:job_catalog;foreignKey:Id" json:"catalogs"`
}

type Company struct {
	ID         *string `gorm:"primary key" json:"id"`
	Name       *string `json:"name"`
	Desc       *string `json:"desc"`
	Employee   *int    `json:"employee"`
	LocationID *string `json:"locationId"`
}

type Location struct {
	ID   *string `gorm:"primary key" json:"id"`
	Name *string `json:"name"`
}

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(&JobEntity{}, &Catalog{}, &Company{}, &Location{})
}
