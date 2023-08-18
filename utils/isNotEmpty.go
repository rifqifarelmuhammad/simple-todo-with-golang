package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsNotEmpty(ctx *gin.Context, data string, dataName string) bool {
	if data == "" {
		ResponseHandler(ctx, HTTPResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: fmt.Sprintf("%s cannot be empty", dataName),
			ResponseStatus:  RESPONSE_STATUS_FAILED,
		})

		return false
	}

	return true
}
