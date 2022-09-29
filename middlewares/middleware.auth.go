package middlewares

import (
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var errorResponse UnathorizatedError

		errorResponse.Status = "Forbidden"
		errorResponse.Code = http.StatusForbidden
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "Auth Need"

		if ctx.GetHeader("Authorization") == "" {
			ctx.JSON(http.StatusForbidden, errorResponse)
			defer ctx.AbortWithStatus(http.StatusForbidden)
		}

		token, err := util.VerifyTokenHeader(ctx, "JWT_SECRET")

		errorResponse.Status = "Unauthorized"
		errorResponse.Code = http.StatusUnauthorized
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "accessToken invalid or expired"

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			// global value result
			ctx.Set("users", token.Claims)
			// return to next method if token is existed
			ctx.Next()
		}

	})
}
