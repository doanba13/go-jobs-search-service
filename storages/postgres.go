package storages

import (
	"fmt"
	"log"
	"os"

	"github.com/doanba13/job-view-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

type Repository struct {
	DB *gorm.DB
}

var Repo Repository

func NewConnection() {
	c := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.Password, c.User, c.DBName, c.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	models.Migration(db)
	Repo = Repository{
		DB: db,
	}
}
