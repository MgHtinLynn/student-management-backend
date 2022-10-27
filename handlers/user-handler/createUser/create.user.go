package handleCreateUser

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/user/createUser"
	"github.com/MgHtinLynn/final-year-project-mcc/service/errorHandler"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
)

type handler struct {
	service createUser.Service
}

func NewHandlerCreateUser(service createUser.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateUserHandler(ctx *gin.Context) {

	var input createUser.InputCreateUser

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

	_, errCreateUser := h.service.CreateUserService(&input)

	switch errCreateUser {

	case "CREATE_USER_EMAIL_CONFLICT_400":
		err := []errorHandler.ErrorMsg{{Field: "Email", Message: "email must be unique"}}
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err)

	case "CREATE_STUDENT_CONFLICT_409":
		util.APIResponse(ctx, "Npm student already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_STUDENT_FAILED_403":
		util.APIResponse(ctx, "Create new student account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Create new student account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
