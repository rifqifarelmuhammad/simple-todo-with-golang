package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/dto"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/models"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/repository"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/utils"
)

func GetAllTodo(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)
	resposeData := repository.FindTodoByUserId(user.UID)

	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    utils.DEFAULT_RESPONSE_CODE,
		ResponseMessage: utils.DEFAULT_RESPONSE_MESSAGE,
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
	}, resposeData)
}

func CreateTodo(ctx *gin.Context) {
	body := dto.CreateTodoRequest{}
	err := ctx.Bind(&body)
	if err != nil {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Failed to read request body",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})
		return
	}

	isNotEmpty := utils.IsNotEmpty(ctx, body.Title, "Title")
	if !isNotEmpty {
		return
	}

	isNotEmpty = utils.IsNotEmpty(ctx, body.Description, "Description")
	if !isNotEmpty {
		return
	}

	user := utils.GetCurrentUser(ctx)
	todo := repository.CreateTodo(user.UID, body)

	responseData := dto.GeneralTodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		IsCompleted: todo.IsCompleted,
		UpdatedAt:   todo.UpdatedAt,
	}

	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    http.StatusCreated,
		ResponseMessage: "Todo has been created",
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
	}, responseData)
}

func UpdateTodo(ctx *gin.Context) {
	todo := TodoIDValidation(ctx)
	if todo == nil {
		return
	}

	updatedTodo := repository.UpdateIsCompleted(todo)

	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    utils.DEFAULT_RESPONSE_CODE,
		ResponseMessage: "Todo has been updated",
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
	}, updatedTodo)
}

func DeleteTodo(ctx *gin.Context) {
	todo := TodoIDValidation(ctx)
	if todo == nil {
		return
	}

	if todo.IsDeleted {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusMethodNotAllowed,
			ResponseMessage: "Todo was deleted previously",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})
		return
	}

	repository.UpdateIsDeleted(todo)

	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    utils.DEFAULT_RESPONSE_CODE,
		ResponseMessage: "Todo has been deleted",
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
	})
}

func TodoIDValidation(ctx *gin.Context) *models.Todo {
	user := utils.GetCurrentUser(ctx)
	todo := repository.FindTodoById(ctx.Param("todoId"))
	if todo.ID == "" || todo.UserID != user.UID {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Invalid Todo ID",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})
		return nil
	}

	return todo
}
