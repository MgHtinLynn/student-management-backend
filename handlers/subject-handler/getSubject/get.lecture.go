package handlerGetSubject

import (
	"fmt"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/subject/getSubject"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct {
	service getSubject.Service
}

func NewHandlerGetSubject(service getSubject.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetSubjectHandler(ctx *gin.Context) {
	var input getSubject.InputGetSubject

	input.ID, _ = strconv.Atoi(ctx.Param("id"))

	fmt.Println("ID ", input.ID)

	getSubjectById, errResultStudent := h.service.GetSubjectService(&input)

	switch errResultStudent {

	case "RESULT_SUBJECT_NOT_FOUND_404":
		util.APIResponse(ctx, "Subject data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "Subject data successfully", http.StatusOK, http.MethodGet, getSubjectById)
	}
}
