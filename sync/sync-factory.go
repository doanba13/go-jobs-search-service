package sync

import (
	"fmt"

	"gorm.io/gorm"
)

func SyncDataHandler(db *gorm.DB, data string, topic string) {
	switch topic {
	case "update-catalog":
		SaveCatalog(db, data)
	case "update-company":
		SaveCompany(db, data)
	case "update-location":
		SaveLocation(db, data)
	case "update-job":
		SaveJob(db, data)
	default:
		fmt.Println("UnHandler topic")
	}
}
