package handlerGetExamResults

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/exam-result/getExamResults"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service getExamResults.Service
}

func NewHandlerGetExamResults(service getExamResults.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetExamResultsHandler(ctx *gin.Context) {
	getExamResultLists, count, errGetExamResults := h.service.GetExamResultsService(ctx)

	switch errGetExamResults {
	case "RESULT_EXAM_RESULT_NOT_FOUND_404":
		util.APIPaginationResponse(ctx, "ExamResult data is not exists", http.StatusConflict, count, http.MethodPost, nil)

	default:
		util.APIPaginationResponse(ctx, "Results Students data successfully", http.StatusOK, count, http.MethodPost, getExamResultLists)
	}
}
