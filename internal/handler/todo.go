package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/dto"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/repository"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/utils"
)

func GetAllTodo(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)
	resposeData := repository.FindTodoByUserId(user.UID)

	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: utils.DEFAULT_RESPONSE_MESSAGE,
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
		Data:            resposeData,
	})
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

	if body.Title == "" {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Title cannot be empty",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return
	}

	if body.Description == "" {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Description cannot be empty",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

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
		Data:            responseData,
	})
}

func UpdateTodo(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)
	todo := repository.FindTodoById(ctx.Param("todoId"))
	if todo.ID == "" || todo.UserID != user.UID {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Invalid Todo ID",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return
	}

	updatedTodo := repository.UpdateIsCompleted(todo)

	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Todo has been updated",
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
		Data:            updatedTodo,
	})
}

func DeleteTodo(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)
	todo := repository.FindTodoById(ctx.Param("todoId"))
	if todo.ID == "" || todo.UserID != user.UID {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Invalid Todo ID",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return
	}

	if todo.IsDeleted {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: "Todo was deleted previously",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})

		return
	}

	repository.UpdateIsDeleted(todo)

	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "Todo has been deleted",
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
	})
}
