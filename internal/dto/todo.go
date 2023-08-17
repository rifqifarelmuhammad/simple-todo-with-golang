package dto

type (
	CreateTodoRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
)
