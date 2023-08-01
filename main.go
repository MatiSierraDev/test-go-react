package main

import (
	_ "fmt"
	"log"
	"os"

	"github.com/MatiSierraDev/6-react-fiber/models"
	"github.com/MatiSierraDev/6-react-fiber/pkg/config"
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

	app.Get("/test", func(c *fiber.Ctx) error {

		db, err := config.DBconfig()
		if err != nil {
			log.Fatal(err)
		}

		// SELECT users.user_id, user_name, user_email, task_id,
		// task_title, task_description
		// FROM users
		// INNER JOIN tasks
		// on users.user_id = tasks.user_id;
		query := `SELECT users.user_id, user_name, user_email, task_id, task_title, task_description
							FROM users
							INNER JOIN tasks
							on users.user_id = tasks.user_id;`

		// query = `SELECT * FROM users`

		rows, err := db.Query(query)

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		var users []models.User
		// var tasks Task

		for rows.Next() {
			var user models.User
			var task models.Task

			err := rows.Scan(&user.Id, &user.Name, &user.Email, &task.Id, &task.Title, &task.Description)
			if err != nil {
				log.Fatal(err)
			}
			user.Tasks = append(user.Tasks, task)
			users = append(users, user)

		}

		return c.Status(200).JSON(fiber.Map{
			"data": users})
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"name": "Grame",
			"age":  20,
		})
	})
	app.Listen(":" + port)
}
