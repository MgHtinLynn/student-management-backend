package handlerGetUser

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/user/getUser"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct {
	service getUser.Service
}

func NewHandlerGetUser(service getUser.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetUserHandler(ctx *gin.Context) {
	var input getUser.InputGetUser

	input.ID, _ = strconv.Atoi(ctx.Param("id"))

	getUserById, errResultStudent := h.service.GetUserService(&input)

	switch errResultStudent {

	case "RESULT_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "user data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "user data successfully", http.StatusOK, http.MethodGet, getUserById)
	}
}
