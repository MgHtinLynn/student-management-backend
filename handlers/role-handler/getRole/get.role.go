package handlerGetRole

import (
	getRoles "github.com/MgHtinLynn/final-year-project-mcc/controllers/roles"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service getRoles.Service
}

func NewHandlerGetRoles(service getRoles.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetRolesHandler(ctx *gin.Context) {
	getRoleLists, count, errGetRoles := h.service.GetRolesService(ctx)

	switch errGetRoles {
	case "ROLES_NOT_FOUND_404":
		util.APIPaginationResponse(ctx, "Roles is not exists", http.StatusConflict, count, http.MethodGet, nil)
	default:
		util.APIPaginationResponse(ctx, "Roles data successfully", http.StatusOK, count, http.MethodGet, getRoleLists)
	}
}
