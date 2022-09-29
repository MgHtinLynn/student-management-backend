package handlerLogin

import (
	"fmt"
	loginAuth "github.com/MgHtinLynn/final-year-project-mcc/controllers/auth-controllers/login"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type handler struct {
	service loginAuth.Service
}

func NewHandlerLogin(service loginAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {

	var input loginAuth.InputLogin

	fmt.Sprintln(input)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultLogin, errLogin := h.service.LoginService(&input)

	switch errLogin {

	case "LOGIN_NOT_FOUND_404":
		util.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
		return

	case "LOGIN_NOT_ACTIVE_403":
		util.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)
		return

	case "LOGIN_WRONG_PASSWORD_403":
		util.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email}
		accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", 24*60*1)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		util.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken, "id": strconv.Itoa(int(resultLogin.ID)), "email": resultLogin.Email, "name": resultLogin.Name})
		return
	}
}
