package main

import (
	"Jimbo8702/saaspack/internal/api"
	"Jimbo8702/saaspack/internal/db"
	"Jimbo8702/saaspack/internal/logger"
	"Jimbo8702/saaspack/internal/validator"
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
    ErrorHandler: api.ErrorHandler,
}

func main() {
	mongoEndpoint := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	var (
		logger         = logger.NewFMTLogger()
		validator      = validator.NewPlaygroundValidator("params")
		productStore   = db.NewMongoProductStore(client)
		productHandler = api.NewProductHandler(productStore, logger, validator)
		app 		   = fiber.New(config)
		apiv1 		   = app.Group("api/v1")
	)
	
	api.AddProductRoutes(apiv1,  productHandler)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}