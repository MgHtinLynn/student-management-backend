package handlerGetSubjects

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/subject/getSubjects"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service getSubjects.Service
}

func NewHandlerGetSubjects(service getSubjects.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetSubjectsHandler(ctx *gin.Context) {
	getSubjectLists, count, errGetSubjects := h.service.GetSubjectsService(ctx)

	switch errGetSubjects {
	case "Subjects_NOT_FOUND_404":
		util.APIPaginationResponse(ctx, "Subject data is not exists", http.StatusConflict, count, http.MethodPost, nil)

	default:
		util.APIPaginationResponse(ctx, "Results Students data successfully", http.StatusOK, count, http.MethodPost, getSubjectLists)
	}
}
