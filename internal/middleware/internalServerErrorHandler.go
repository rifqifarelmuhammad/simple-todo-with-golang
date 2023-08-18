package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InternalServerErrorHTTPResponse struct {
	ResponseCode    int
	ResponseMessage string
	ResponseStatus  string
}

func InternalServerErrorHandler(ctx *gin.Context, err any) {
	response := InternalServerErrorHTTPResponse{
		ResponseCode:    http.StatusInternalServerError,
		ResponseMessage: "Request failed with status code 500",
		ResponseStatus:  "FAILED",
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
}
