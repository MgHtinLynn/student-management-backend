package handleCreateExamResult

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/exam-result/createExamResult"
	"github.com/MgHtinLynn/final-year-project-mcc/service/errorHandler"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
)

type handler struct {
	service createExamResult.Service
}

func NewHandlerCreateExamResult(service createExamResult.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateExamResultHandler(ctx *gin.Context) {

	var input createExamResult.InputCreateExamResult

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

	_, errCreateExamResult := h.service.CreateExamResultService(&input)

	switch errCreateExamResult {

	case "CREATE_ExamResult_CONFLICT_409":
		util.APIResponse(ctx, "ExamResult already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_ExamResult_FAILED_403":
		util.APIResponse(ctx, "Create ExamResult account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Create ExamResult account successfully", http.StatusCreated, http.MethodPost, nil)
		return
	}
}
