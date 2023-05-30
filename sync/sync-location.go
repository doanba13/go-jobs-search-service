package sync

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type Location struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

func SaveLocation(db *gorm.DB, data string) {
	c := Location{}
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
