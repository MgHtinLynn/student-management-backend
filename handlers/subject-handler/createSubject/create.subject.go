package handleCreateSubject

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/subject/createSubject"
	"github.com/MgHtinLynn/final-year-project-mcc/service/errorHandler"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
)

type handler struct {
	service createSubject.Service
}

func NewHandlerCreateSubject(service createSubject.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateSubjectHandler(ctx *gin.Context) {

	var input createSubject.InputCreateSubject

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

	_, errCreateSubject := h.service.CreateSubjectService(&input)

	switch errCreateSubject {

	case "CREATE_Subject_EMAIL_CONFLICT_400":
		err := []errorHandler.ErrorMsg{{Field: "Email", Message: "email must be unique"}}
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err)

	case "CREATE_STUDENT_CONFLICT_409":
		util.APIResponse(ctx, "Subject already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_Subject_FAILED_403":
		util.APIResponse(ctx, "Create Subject account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Create Subject account successfully", http.StatusCreated, http.MethodPost, nil)
		return
	}
}
