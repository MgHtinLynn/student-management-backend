package handleUpdateSubject

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/subject/updateSubject"
	"github.com/MgHtinLynn/final-year-project-mcc/service/errorHandler"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type handler struct {
	service updateSubject.Service
}

func NewHandlerUpdateSubject(service updateSubject.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateSubjectHandler(ctx *gin.Context) {

	var input updateSubject.InputUpdateSubject
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

	getSubject, errUpdateSubject := h.service.UpdateSubjectService(&input)

	switch errUpdateSubject {

	case "UPDATE_Subject_EMAIL_CONFLICT_400":
		err := []errorHandler.ErrorMsg{{Field: "email", Message: "email must be unique"}}
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err)

	case "UPDATE_Subject_NOT_FOUND_404":
		util.APIResponse(ctx, "Subject data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)

	case "UPDATE_Subject_FAILED_403":
		util.APIResponse(ctx, "Update Subject data failed", http.StatusForbidden, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Update Subject data successfully", http.StatusOK, http.MethodPost, getSubject)
	}
}
