package main

import (
	_ "fmt"
	"log"
	"os"

	"github.com/MatiSierraDev/6-react-fiber/pkg/config"
	"github.com/MatiSierraDev/6-react-fiber/pkg/controllers"
	"github.com/MatiSierraDev/6-react-fiber/pkg/models"
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

	app.Get("/test", apiRoute.GetAlls)
	app.Post("/test", func(ctx *fiber.Ctx) error {

		DB, err := config.DBconfig()

		if err != nil {
			log.Fatal(err)
		}

		var newUser models.User

		newUser.Name = "Chinchu"
		newUser.Email = "Chinchu@gmail.com"

		query := `INSERT INTO users(user_name, user_email)
						VALUES ($1,$2) RETURNING *`

		// result, err := DB.Exec(query, newUser.Name, newUser.Email)

		//guardo(&scan) los datos que me devuelve la DB en mi variable
		err = DB.QueryRow(query, newUser.Name, newUser.Email).Scan(&newUser.Id, &newUser.Name, &newUser.Email)

		if err != nil {
			log.Fatal(err)
		}

		// data, err := result.RowsAffected()

		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(&newUser)

	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"name": "Grame",
			"age":  20,
		})
	})
	app.Listen(":" + port)
}
