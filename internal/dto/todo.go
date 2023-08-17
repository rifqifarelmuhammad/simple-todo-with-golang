package dto

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateTodoRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	CreateTodoResponse struct {
		ID          uuid.UUID `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		IsCompleted bool      `json:"is_completed"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
