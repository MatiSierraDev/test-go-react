package main

import (
	_ "fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	//Midleware
	app.Use(cors.New())

	//Server files
	app.Static("/", "./client/dist")

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"name": "Grame",
			"age":  20,
		})
	})
	app.Listen(":" + port)
}
