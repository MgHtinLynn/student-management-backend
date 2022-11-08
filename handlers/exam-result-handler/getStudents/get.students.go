package handlerGetStudents

import (
	getStudents "github.com/MgHtinLynn/final-year-project-mcc/controllers/student"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service getStudents.Service
}

func NewHandlerGetStudents(service getStudents.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetStudentsHandler(ctx *gin.Context) {
	getStudentList, count, errGetStudents := h.service.GetStudentsService(ctx)

	switch errGetStudents {
	case "STUDENTS_NOT_FOUND_404":
		util.APIPaginationResponse(ctx, "Students data is not exists", http.StatusConflict, count, http.MethodPost, nil)

	default:
		util.APIPaginationResponse(ctx, "Results Students data successfully", http.StatusOK, count, http.MethodPost, getStudentList)
	}
}
