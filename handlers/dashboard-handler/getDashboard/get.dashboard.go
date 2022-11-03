package handlerGetDashboard

import (
	getDashboard "github.com/MgHtinLynn/final-year-project-mcc/controllers/dashboard"
	//model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct {
	service getDashboard.Service
}

func NewHandlerGetDashboard(service getDashboard.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetDashboardHandler(ctx *gin.Context) {

	var input getDashboard.InputGetUser

	input.ID, _ = strconv.Atoi(ctx.Param("id"))

	user, count, activeCount, errGetDashboard := h.service.GetDashboardService(&input)

	data := util.DataDashboard{
		ActiveCount: activeCount,
		Count:       count,
		User:        user,
	}

	switch errGetDashboard {
	case "DASHBOARD_NOT_FOUND_404":
		util.APIResponse(ctx, "something wrong", http.StatusForbidden, http.MethodPost, nil)
	default:
		util.APIDashboardResponse(ctx, "Success", http.StatusOK, http.MethodPost, data)
	}
}
