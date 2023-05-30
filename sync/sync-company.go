package sync

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type Company struct {
	Id         string `json:"_id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Employee   int    `json:"employee"`
	LocationID string `json:"locationId"`
}

func SaveCompany(db *gorm.DB, data string) {
	c := Company{}
	err := json.Unmarshal([]byte(data), &c)

	fmt.Println(c)

	if err != nil {
		fmt.Println(err)
	}

	err = db.Create(&c).Error

	if err != nil {
		fmt.Println(err)
	}
}
