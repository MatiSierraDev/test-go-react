package controllers

import (
	"log"

	"github.com/MatiSierraDev/6-react-fiber/pkg/config"
	"github.com/MatiSierraDev/6-react-fiber/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type Api struct{}

func (a *Api) GetAlls(ctx *fiber.Ctx) error {
	db, err := config.DBconfig()
	if err != nil {
		log.Fatal(err)
	}

	query := `SELECT users.user_id, user_name, user_email, task_id, task_title, task_description
						FROM users
						LEFT JOIN tasks
						on users.user_id = tasks.user_id;`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var users []models.User

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

	return ctx.Status(200).JSON(fiber.Map{
		"data": users})
}
func (a *Api) CreateUser(ctx *fiber.Ctx) error {

	DB, err := config.DBconfig()

	if err != nil {
		log.Fatal(err)
	}

	var newUser models.User

	newUser.Name = "Messi"
	newUser.Email = "Messi@gmail.com"

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
}
