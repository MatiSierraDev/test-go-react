package models

import (
	_ "fmt"

	"github.com/google/uuid"
)

type User struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Tasks []Task    `json:"tasks"`
}
