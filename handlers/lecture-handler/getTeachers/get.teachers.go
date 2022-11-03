package handlerGetLectures

import (
	getTeachers "github.com/MgHtinLynn/final-year-project-mcc/controllers/teacher"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service getTeachers.Service
}

func NewHandlerGetTeachers(service getTeachers.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetTeachersHandler(ctx *gin.Context) {
	getTeacherList, count, errGetTeachers := h.service.GetTeachersService(ctx)

	switch errGetTeachers {
	case "TEACHERS_NOT_FOUND_404":
		util.APIPaginationResponse(ctx, "Teachers data is not exists", http.StatusConflict, count, http.MethodPost, nil)

	default:
		util.APIPaginationResponse(ctx, "Results Teachers data successfully", http.StatusOK, count, http.MethodPost, getTeacherList)
	}
}
