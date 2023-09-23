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
		//logger
		logger          = logger.NewFMTLogger()
		//validator
		validator       = validator.NewPlaygroundValidator("params")
		//stores
		productStore    = db.NewMongoProductStore(client)
		categoryStore   = db.NewMongoCategoryStore(client)
		bookingStore 	= db.NewMongoBookingStore(client)
		profileStore 	= db.NewMongoProfileStore(client)
		//handlers
		bookingHandler 	= api.NewBookingHandler(bookingStore, logger, validator)
		profileHandler 	= api.NewProfileHandler(profileStore, logger, validator)
		productHandler  = api.NewProductHandler(productStore, logger, validator)
		categoryHandler = api.NewCategoryHandler(categoryStore, logger, validator)
		//app & app groups
		app 		    = fiber.New(config)
		apiv1 		    = app.Group("api/v1")
	)
	
	api.AddProductRoutes(apiv1,  productHandler)
	api.AddCategoryRoutes(apiv1, categoryHandler)
	api.AddBookingRoutes(apiv1, bookingHandler)
	api.AddProfileRoutes(apiv1, profileHandler)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}