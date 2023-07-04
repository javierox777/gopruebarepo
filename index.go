package main

import (
	modeluser "backend/model"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://locahost:27018/gouser"))
	if err != nil {
		panic(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		var users []modeluser.User
		call := client.Database("gouser").Collection("user")
		result, err := call.Find(context.TODO(), bson.M{})
		if err != nil {
			panic(err)
		}
		fmt.Println("hola mundo")
		for result.Next(context.TODO()) {
			var user modeluser.User
			result.Decode(&user)
			users = append(users, user)
		}

		return c.JSON(&fiber.Map{
			"data": users,
		})
	})

	app.Listen(":3000")
}
