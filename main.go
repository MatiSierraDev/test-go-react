package main

import (
	_ "fmt"
	"os"

	"github.com/MatiSierraDev/6-react-fiber/pkg/config"
	"github.com/MatiSierraDev/6-react-fiber/pkg/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/google/uuid"
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

	config.DBconfig()

	apiRoute := controllers.Api{}

	app.Get("/api/users", apiRoute.GetAlls)
	app.Post("/api/users", apiRoute.CreateUser)

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"name": "Grame",
			"age":  20,
		})
	})
	app.Listen(":" + port)
}
