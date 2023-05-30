package sync

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type Catalog struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

func SaveCatalog(db *gorm.DB, data string) {
	c := Catalog{}
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
