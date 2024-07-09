package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/imakhlaq/hotelreservation/api/apiv1/handlers"
	"github.com/imakhlaq/hotelreservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUri    = "mongodb://localhost:27017"
	dbName   = "hotel-reservation"
	userColl = "user"
)

func main() {
	//connecting to mongodb
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}
	coll := client.Database(dbName).Collection(userColl)

	user := types.User{
		FirstName: "Akhlaq",
		LastName:  "Ahmad",
	}
	res, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	var akhlaq types.User
	if err := coll.FindOne(context.Background(), bson.M{}).Decode(&akhlaq); err != nil {
		log.Fatal(err)
	}
	fmt.Println(akhlaq)

	//go run main.go --listenAddr :9000  => if u want to give address
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/users", handlers.HandleUsers)
	apiV1.Get("/user/:id", handlers.HandleUser)

	app.Listen(*listenAddr)
}
