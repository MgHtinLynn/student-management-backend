package handlerGetLectures

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/lecture/getLectures"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service getLectures.Service
}

func NewHandlerGetLectures(service getLectures.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetLecturesHandler(ctx *gin.Context) {
	getLectureLists, count, errGetLectures := h.service.GetLecturesService(ctx)

	switch errGetLectures {
	case "Lectures_NOT_FOUND_404":
		util.APIPaginationResponse(ctx, "Lecture data is not exists", http.StatusConflict, count, http.MethodPost, nil)

	default:
		util.APIPaginationResponse(ctx, "Results Students data successfully", http.StatusOK, count, http.MethodPost, getLectureLists)
	}
}
