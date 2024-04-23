package models

import (
	"github.com/google/uuid"
)

type TodoItem struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"isCompleted"`
}
