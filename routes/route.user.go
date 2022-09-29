package routes

import (
	getUsers "github.com/MgHtinLynn/final-year-project-mcc/controllers/auth-controllers/users"
	handleResultUsers "github.com/MgHtinLynn/final-year-project-mcc/handlers/user-handler/getUser"
	"github.com/MgHtinLynn/final-year-project-mcc/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {

	usersRepository := getUsers.NewRepositoryUser(db)
	usersService := getUsers.NewServiceUsers(usersRepository)
	usersHandler := handleResultUsers.NewHandlerGetUsers(usersService)
	/**
	@description All User Route
	*/
	groupRoute := route.Group("/api/v1").Use(middlewares.Auth())
	groupRoute.GET("/users", usersHandler.GetUsersHandler)

}
