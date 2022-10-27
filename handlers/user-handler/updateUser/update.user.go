package handleUpdateUser

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/user/updateUser"
	"github.com/MgHtinLynn/final-year-project-mcc/service/errorHandler"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type handler struct {
	service updateUser.Service
}

func NewHandlerUpdateUser(service updateUser.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateUserHandler(ctx *gin.Context) {

	var input updateUser.InputUpdateUser
	input.ID, _ = strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]errorHandler.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = errorHandler.ErrorMsg{Field: fe.Field(), Message: util.GetErrorMsg(fe)}
			}
			util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, out)
		}
		return
	}

	getUser, errUpdateUser := h.service.UpdateUserService(&input)

	switch errUpdateUser {

	case "UPDATE_USER_EMAIL_CONFLICT_400":
		err := []errorHandler.ErrorMsg{{Field: "email", Message: "email must be unique"}}
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err)

	case "UPDATE_USER_NOT_FOUND_404":
		util.APIResponse(ctx, "user data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)

	case "UPDATE_USER_FAILED_403":
		util.APIResponse(ctx, "Update user data failed", http.StatusForbidden, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Update user data successfully", http.StatusOK, http.MethodPost, getUser)
	}
}
