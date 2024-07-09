package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/imakhlaq/hotelreservation/api/apiv1/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUri  = "mongodb://localhost:27017"
	dbName = "hotel-reservation"
)

func main() {
	//connecting to mongodb
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	//go run main.go --listenAddr :9000  => if u want to give address
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()

	apiV1 := app.Group("/api/v1")
	apiV1.Get("/users", handlers.HandleUsers)
	apiV1.Get("/user/:id", handlers.HandleUser)

	app.Listen(*listenAddr)
}
