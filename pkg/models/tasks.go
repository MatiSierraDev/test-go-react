package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Task struct {
	Id *uuid.UUID `json:"id"`
	// Title       *string    `json:"title"`
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
}
