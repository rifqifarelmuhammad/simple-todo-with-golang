package dto

import (
	"time"
)

type (
	CreateTodoRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	GeneralTodoResponse struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		IsCompleted bool      `json:"is_completed"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	UpdateTodoResponse struct {
		IsCompleted bool      `json:"is_completed"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
