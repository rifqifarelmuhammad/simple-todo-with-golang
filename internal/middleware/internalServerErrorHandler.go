package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rifqifarelmuhammad/simple-todo-with-golang/utils"
)

type PanicHTTPResponse struct {
	ResponseCode    int
	ResponseMessage string
	ResponseStatus  string
}

func PanicHandler(ctx *gin.Context, err any) {
	var response PanicHTTPResponse
	if ctx.Writer.Status() == 403 {
		utils.ResponseHandler(ctx, utils.HTTPResponse{
			ResponseCode:    http.StatusForbidden,
			ResponseMessage: "Forbidden resource",
			ResponseStatus:  "FAILED",
		})
	} else {
		response = PanicHTTPResponse{
			ResponseCode:    http.StatusInternalServerError,
			ResponseMessage: "Request failed with status code 500",
			ResponseStatus:  "FAILED",
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
	}
}
