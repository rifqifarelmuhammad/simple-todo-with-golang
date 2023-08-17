package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/utils"
)

func GetTodo(ctx *gin.Context) {
	utils.ResponseHandler(ctx, utils.HTTPResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: utils.DEFAULT_RESPONSE_MESSAGE,
		ResponseStatus:  utils.RESPONSE_STATUS_SUCCESS,
	})
}
