package repository

import (
	"github.com/google/uuid"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/constant"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/database"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/dto"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/models"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/log"
)

func FindTodoByUserId(userId uuid.UUID) *[]dto.GeneralTodoResponse {
	todos := &[]dto.GeneralTodoResponse{}
	result := database.GetInstance().Raw("SELECT id, title, description, is_completed, updated_at FROM todos WHERE user_id = ?", userId).Scan(todos)
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "todo[FindTodoByUserId]: Error query db on database.GetInstance().Find")
		panic(result.Error)
	}

	return todos
}

func CreateTodo(uid uuid.UUID, data dto.CreateTodoRequest) *models.Todo {
	todo := &models.Todo{
		ID:          uuid.New(),
		Title:       data.Title,
		Description: data.Description,
		UserID:      uid,
	}

	result := database.GetInstance().Create(todo)
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "todo[CreateTodo]: Error query db on database.GetInstance().Create")
		panic(result.Error)
	}

	return todo
}
