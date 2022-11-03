package routes

import (
	getDashboard "github.com/MgHtinLynn/final-year-project-mcc/controllers/dashboard"
	handlerGetDashboard "github.com/MgHtinLynn/final-year-project-mcc/handlers/dashboard-handler/getDashboard"
	"github.com/MgHtinLynn/final-year-project-mcc/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitDashboardRoutes(db *gorm.DB, route *gin.Engine) {

	dashboardRepository := getDashboard.NewRepositoryDashboard(db)
	dashboardService := getDashboard.NewServiceDashboard(dashboardRepository)
	dashboardHandler := handlerGetDashboard.NewHandlerGetDashboard(dashboardService)

	/**
	@description All Role Route
	*/
	groupRoute := route.Group("/api/v1").Use(middlewares.Auth())
	groupRoute.GET("/dashboard/:id", dashboardHandler.GetDashboardHandler)
}
