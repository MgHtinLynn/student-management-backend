package handlerGetExamResult

import (
	"fmt"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/exam-result/getExamResult"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct {
	service getExamResult.Service
}

func NewHandlerGetExamResult(service getExamResult.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetExamResultHandler(ctx *gin.Context) {
	var input getExamResult.InputGetExamResult

	input.ID, _ = strconv.Atoi(ctx.Param("id"))

	fmt.Println("ID ", input.ID)

	getExamResultById, errResultStudent := h.service.GetExamResultService(&input)

	switch errResultStudent {

	case "RESULT_SUBJECT_NOT_FOUND_404":
		util.APIResponse(ctx, "ExamResult data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "ExamResult data successfully", http.StatusOK, http.MethodGet, getExamResultById)
	}
}
