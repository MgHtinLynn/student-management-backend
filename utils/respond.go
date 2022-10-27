package utils

import "github.com/gin-gonic/gin"

type PaginationResponses struct {
	StatusCode int         `json:"status_code"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Total      *int64      `json:"total" bson:"total"`
	Data       interface{} `json:"data"`
}

type Responses struct {
	StatusCode int         `json:"status_code"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Total      *int64      `json:"total" bson:"total"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Errors     interface{} `json:"errors"`
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Method string, Data interface{}) {

	jsonResponse := Responses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}

}

func APIPaginationResponse(ctx *gin.Context, Message string, StatusCode int, Total *int64, Method string, Data interface{}) {

	jsonResponse := PaginationResponses{
		StatusCode: StatusCode,
		Method:     Method,
		Total:      Total,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}

}

func ValidatorErrorResponse(ctx *gin.Context, StatusCode int, Method string, Errors interface{}) {
	errResponse := ErrorResponse{
		StatusCode: StatusCode,
		Method:     Method,
		Errors:     Errors,
	}

	ctx.JSON(StatusCode, errResponse)
	defer ctx.AbortWithStatus(StatusCode)
}
