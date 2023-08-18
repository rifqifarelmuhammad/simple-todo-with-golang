package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/utils"
)

func PanicHandler(ctx *gin.Context, err any) {
	if ctx.Writer.Status() == http.StatusForbidden {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusForbidden,
			ResponseMessage: "Forbidden resource",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		})
	} else {
		response := utils.HTTPResponse{
			ResponseCode:    http.StatusInternalServerError,
			ResponseMessage: "Request failed with status code 500",
			ResponseStatus:  utils.RESPONSE_STATUS_FAILED,
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
	}
}
