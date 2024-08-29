package models

import "github.com/google/uuid"

type Task struct {
	BaseModel
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	UserID      uuid.UUID `json:"userId"`
}
