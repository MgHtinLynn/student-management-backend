package handlerGetTranscripts

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/transcript/getTranscripts"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service getTranscripts.Service
}

func NewHandlerGetTranscripts(service getTranscripts.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetTranscriptsHandler(ctx *gin.Context) {
	getTranscriptLists, count, errGetTranscripts := h.service.GetTranscriptsService(ctx)

	switch errGetTranscripts {
	case "RESULT_EXAM_RESULT_NOT_FOUND_404":
		util.APIPaginationResponse(ctx, "Transcript data is not exists", http.StatusConflict, count, http.MethodPost, nil)

	default:
		util.APIPaginationResponse(ctx, "Results Transcript data successfully", http.StatusOK, count, http.MethodPost, getTranscriptLists)
	}
}
