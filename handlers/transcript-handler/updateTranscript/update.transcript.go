package handleUpdateTranscript

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/transcript/updateTranscript"
	"github.com/MgHtinLynn/final-year-project-mcc/service/errorHandler"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type handler struct {
	service updateTranscript.Service
}

func NewHandlerUpdateTranscript(service updateTranscript.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateTranscriptHandler(ctx *gin.Context) {

	var input updateTranscript.InputUpdateTranscript
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

	getTranscript, errUpdateTranscript := h.service.UpdateTranscriptService(&input)

	switch errUpdateTranscript {

	case "UPDATE_Transcript_EMAIL_400":
		err := []errorHandler.ErrorMsg{{Field: "email", Message: "email must be unique"}}
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err)

	case "UPDATE_Transcript_NOT_FOUND_404":
		util.APIResponse(ctx, "Transcript data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)

	case "UPDATE_Transcript_FAILED_403":
		util.APIResponse(ctx, "Update Transcript data failed", http.StatusForbidden, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Update Transcript data successfully", http.StatusOK, http.MethodPost, getTranscript)
	}
}
