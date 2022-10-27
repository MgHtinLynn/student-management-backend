package routes

import (
	getRoles "github.com/MgHtinLynn/final-year-project-mcc/controllers/roles"
	handlerGetRole "github.com/MgHtinLynn/final-year-project-mcc/handlers/role-handler/getRole"
	"github.com/MgHtinLynn/final-year-project-mcc/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoleRoutes(db *gorm.DB, route *gin.Engine) {

	rolesRepository := getRoles.NewRepositoryRole(db)
	rolesService := getRoles.NewServiceRoles(rolesRepository)
	rolesHandler := handlerGetRole.NewHandlerGetRoles(rolesService)

	/**
	@description All Role Route
	*/
	groupRoute := route.Group("/api/v1").Use(middlewares.Auth())
	groupRoute.GET("/roles", rolesHandler.GetRolesHandler)
}
