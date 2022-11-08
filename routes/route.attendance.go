package routes

import (
	getAttendances "github.com/MgHtinLynn/final-year-project-mcc/controllers/attendance"
	handlerGetAttendance "github.com/MgHtinLynn/final-year-project-mcc/handlers/attendance-handler/getAttendance"
	"github.com/MgHtinLynn/final-year-project-mcc/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAttendanceRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All ExamResult Route
	*/

	attendancesRepository := getAttendances.NewRepositoryAttendance(db)
	service := getAttendances.NewServiceAttendances(attendancesRepository)
	attendancesHandler := handlerGetAttendance.NewHandlerGetAttendances(service)

	//
	groupRoute := route.Group("/api/v1").Use(middlewares.Auth())
	groupRoute.GET("/attendances", attendancesHandler.GetAttendancesHandler)

}
