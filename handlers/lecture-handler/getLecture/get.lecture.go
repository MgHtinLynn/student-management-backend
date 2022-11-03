package handlerGetLecture

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/lecture/getLecture"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct {
	service getLecture.Service
}

func NewHandlerGetLecture(service getLecture.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetLectureHandler(ctx *gin.Context) {
	var input getLecture.InputGetLecture

	input.ID, _ = strconv.Atoi(ctx.Param("id"))

	getLectureById, errResultStudent := h.service.GetLectureService(&input)

	switch errResultStudent {

	case "RESULT_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "lecture data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "lecture data successfully", http.StatusOK, http.MethodGet, getLectureById)
	}
}
