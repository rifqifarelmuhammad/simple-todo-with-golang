package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPResponse struct {
	ResponseCode    int
	ResponseMessage string
	ResponseStatus  string
}

const (
	DEFAULT_RESPONSE_CODE    = http.StatusOK
	DEFAULT_RESPONSE_MESSAGE = "Data retrieved successfully"
	RESPONSE_STATUS_SUCCESS  = "SUCCESS"
	RESPONSE_STATUS_FAILED   = "FAILED"
)

func ResponseHandler(ctx *gin.Context, httpResponse HTTPResponse, data ...interface{}) {
	responseBody := make(map[string]interface{})
	responseBody["responseCode"] = httpResponse.ResponseCode
	responseBody["responseMessage"] = httpResponse.ResponseMessage
	responseBody["responseStatus"] = httpResponse.ResponseStatus

	if data != nil {
		if len(data) == 1 {
			responseBody["data"] = data[0]
		} else {
			responseBody["data"] = data
		}
	}

	ctx.JSON(httpResponse.ResponseCode, responseBody)
}
