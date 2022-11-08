package handlerGetAttendance

import (
	getAttendances "github.com/MgHtinLynn/final-year-project-mcc/controllers/attendance"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service getAttendances.Service
}

func NewHandlerGetAttendances(service getAttendances.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAttendancesHandler(ctx *gin.Context) {
	getAttendanceLists, count, errGetAttendances := h.service.GetAttendancesService(ctx)

	switch errGetAttendances {
	case "Attendances_NOT_FOUND_404":
		util.APIPaginationResponse(ctx, "Attendances is not exists", http.StatusConflict, count, http.MethodGet, nil)
	default:
		util.APIPaginationResponse(ctx, "Attendances data successfully", http.StatusOK, count, http.MethodGet, getAttendanceLists)
	}
}
