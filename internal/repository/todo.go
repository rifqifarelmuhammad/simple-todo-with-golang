package repository

import (
	"github.com/google/uuid"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/constant"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/database"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/dto"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/models"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/log"
)

func FindTodoById(id string) *models.Todo {
	// fmt.Println(id)
	todo := &models.Todo{}
	result := database.GetInstance().Find(todo, models.Todo{ID: id})
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "todo[FindTodoById]: Error query db on database.GetInstance().Find(args)")
		panic(result.Error)
	}

	return todo
}

func FindTodoByUserId(userId string) *[]dto.GeneralTodoResponse {
	todos := &[]dto.GeneralTodoResponse{}
	result := database.GetInstance().Raw("SELECT id, title, description, is_completed, updated_at FROM todos WHERE user_id = ?", userId).Scan(todos)
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "todo[FindTodoByUserId]: Error query db on database.GetInstance().Raw(args).Scan")
		panic(result.Error)
	}

	return todos
}

func CreateTodo(uid string, data dto.CreateTodoRequest) *models.Todo {
	todo := &models.Todo{
		ID:          (uuid.New()).String(),
		Title:       data.Title,
		Description: data.Description,
		UserID:      uid,
	}

	result := database.GetInstance().Create(todo)
	if result.Error != nil {
		log.Error(constant.TAG_REPOSITORY, result, result.Error, "todo[CreateTodo]: Error query db on database.GetInstance().Create(args)")
		panic(result.Error)
	}

	return todo
}
