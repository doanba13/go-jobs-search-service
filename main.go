package main

import (
	"context"
	"fmt"
	"log"

	"github.com/doanba13/job-view-service/router"
	"github.com/doanba13/job-view-service/storages"
	"github.com/doanba13/job-view-service/sync"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	kafka "github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/job-view")

	api.Post("/job", router.CreateJob)
	api.Get("/job", router.GetAllJob)
	api.Get("/job/:id", router.DetailJob)
	api.Get("/search-job", router.SearchJob)
	api.Get("/recently-job", router.GetRecentlyJob)
	api.Get("/high-salary-job", router.GetHighSalaryJob)
	api.Get("/hot-job", router.GetJobByView)

	api.Get("/company", router.GetAllCompany)
	api.Get("/company/:id", router.GetJobByCompanyId)

	api.Get("/location", router.GetAllLocation)

	api.Get("/catalog/:id", router.GetAllCatalogJob)
	api.Get("/catalog", router.GetAllCatalog)

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	// Connect to DB
	storages.NewConnection()

	// Start kafka client
	kafkaConsumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "job-view",
		StartOffset: kafka.LastOffset,
		GroupTopics: []string{"update-job", "update-catalog", "update-company", "update-location"},
	})
	go func(db *gorm.DB) {
		fmt.Println("start consuming ... !!")

		for {
			msg, err := kafkaConsumer.ReadMessage(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("topic: ", msg.Topic)
			fmt.Println("value: ", string(msg.Value))
			sync.SyncDataHandler(db, string(msg.Value), msg.Topic)
		}
	}(storages.Repo.DB)

	// Setup fiber and router
	app := fiber.New()
	SetupRoutes(app)
	app.Listen(":8802")

	// Grateful shutdown
	// signalChan := make(chan os.Signal, 1)
	// signal.Notify(signalChan, os.Interrupt)
	// signal.Notify(signalChan, os.Kill)

	// sig := <-signalChan
	// log.Println("Recieved terminal, grateful shutdown in 10s", sig)

	// //timeOutContext
	// _, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// defer cancel()
	defer kafkaConsumer.Close()

	// Disconnect all connection here:
	// s.Shutdown(timeOutContext)
}
