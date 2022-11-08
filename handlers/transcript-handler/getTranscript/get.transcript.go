package handlerGetTranscript

import (
	"fmt"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/transcript/getTranscript"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct {
	service getTranscript.Service
}

func NewHandlerGetTranscript(service getTranscript.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetTranscriptHandler(ctx *gin.Context) {
	var input getTranscript.InputGetTranscript

	input.ID, _ = strconv.Atoi(ctx.Param("id"))

	fmt.Println("ID ", input.ID)

	getTranscriptById, errResultStudent := h.service.GetTranscriptService(&input)

	switch errResultStudent {

	case "RESULT_Transcript_NOT_FOUND_404":
		util.APIResponse(ctx, "Transcript data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "Transcript data successfully", http.StatusOK, http.MethodGet, getTranscriptById)
	}
}
