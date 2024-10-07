package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	handlers "github.com/imakhlaq/hotelreservation/api/apiv1"
	"github.com/imakhlaq/hotelreservation/db"
	"github.com/imakhlaq/hotelreservation/error"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUri = "mongodb://localhost:27017"
)

func main() {
	//go run main.go --listenAddr :9000  => if u want to give address
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	//connecting to mongodb
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(error.Config)
	apiV1 := app.Group("/api/v1")

	//initializing handler with db
	userHandler := handlers.NewUserHandler(db.NewMongoUserStore(client))
	apiV1.Post("/user", userHandler.HandlePostUser)
	apiV1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiV1.Get("/users", userHandler.HandleUsers)
	apiV1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*listenAddr)
}
