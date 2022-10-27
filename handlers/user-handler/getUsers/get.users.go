package handlerGetUsers

import (
	getUsers "github.com/MgHtinLynn/final-year-project-mcc/controllers/auth-controllers/users"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service getUsers.Service
}

func NewHandlerGetUsers(service getUsers.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetUsersHandler(ctx *gin.Context) {
	getUserLists, count, errGetUsers := h.service.GetUsersService(ctx)

	switch errGetUsers {
	case "USERS_NOT_FOUND_404":
		util.APIPaginationResponse(ctx, "User data is not exists", http.StatusConflict, count, http.MethodPost, nil)

	default:
		util.APIPaginationResponse(ctx, "Results Students data successfully", http.StatusOK, count, http.MethodPost, getUserLists)
	}
}
