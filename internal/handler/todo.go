package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/dto"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/internal/repository"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/utils"
)

func GetTodo(ctx *gin.Context) {
	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: utils.DEFAULT_RESPONSE_MESSAGE,
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
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
	todo, err := repository.CreateTodo(user.UID, body)
	if err != nil {
		return
	}

	responseData := dto.CreateTodoResponse{
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
