package utils

import "github.com/gin-gonic/gin"

type HTTPResponse struct {
	ResponseCode    int
	ResponseMessage string
	ResponseStatus  string
	Data            interface{}
}

const (
	DEFAULT_RESPONSE_MESSAGE = "Data retrieved successfully"
	RESPONSE_STATUS_SUCCESS  = "SUCCESS"
	RESPONSE_STATUS_FAILED   = "FAILED"
)

func ResponseHandler(ctx *gin.Context, resData HTTPResponse) {
	responseBody := make(map[string]interface{})
	responseBody["responseCode"] = resData.ResponseCode
	responseBody["responseMessage"] = resData.ResponseMessage
	responseBody["responseStatus"] = resData.ResponseStatus

	if resData.Data != nil {
		responseBody["data"] = resData.Data
	}

	ctx.JSON(resData.ResponseCode, responseBody)
}
