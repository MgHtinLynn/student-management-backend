package routes

import (
	getUsers "github.com/MgHtinLynn/final-year-project-mcc/controllers/auth-controllers/users"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/user/createUser"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/user/getUser"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/user/updatePassword"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/user/updateUser"
	handleCreateUser "github.com/MgHtinLynn/final-year-project-mcc/handlers/user-handler/createUser"
	handleResultUser "github.com/MgHtinLynn/final-year-project-mcc/handlers/user-handler/getUser"
	handleResultUsers "github.com/MgHtinLynn/final-year-project-mcc/handlers/user-handler/getUsers"
	handleUpdatePassword "github.com/MgHtinLynn/final-year-project-mcc/handlers/user-handler/updatePassword"
	handleUpdateUser "github.com/MgHtinLynn/final-year-project-mcc/handlers/user-handler/updateUser"
	"github.com/MgHtinLynn/final-year-project-mcc/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All User Route
	*/

	usersRepository := getUsers.NewRepositoryUser(db)
	usersService := getUsers.NewServiceUsers(usersRepository)
	usersHandler := handleResultUsers.NewHandlerGetUsers(usersService)

	userRepository := getUser.NewRepositoryUser(db)
	userService := getUser.NewServiceUser(userRepository)
	userHandler := handleResultUser.NewHandlerGetUser(userService)

	createUserRepository := createUser.NewRepositoryCreate(db)
	createUserService := createUser.NewServiceCreate(createUserRepository)
	createUserHandler := handleCreateUser.NewHandlerCreateUser(createUserService)

	updateUserRepository := updateUser.NewRepositoryUpdate(db)
	updateUserService := updateUser.NewServiceUpdate(updateUserRepository)
	updateUserHandler := handleUpdateUser.NewHandlerUpdateUser(updateUserService)

	updatePasswordRepository := updatePassword.NewRepositoryUpdate(db)
	updatePasswordService := updatePassword.NewServiceUpdate(updatePasswordRepository)
	updatePasswordHandler := handleUpdatePassword.NewHandlerUpdatePassword(updatePasswordService)

	groupRouteCustom := route.Group("/api/v1")
	groupRouteCustom.GET("/userLists", usersHandler.GetUsersHandler)
	groupRouteCustom.PATCH("/changePassword/:id", updatePasswordHandler.UpdatePasswordHandler)

	groupRoute := route.Group("/api/v1").Use(middlewares.Auth())
	groupRoute.GET("/users", usersHandler.GetUsersHandler)
	groupRoute.POST("/users", createUserHandler.CreateUserHandler)
	groupRoute.GET("/users/:id", userHandler.GetUserHandler)
	groupRoute.PUT("/users/:id", updateUserHandler.UpdateUserHandler)

}
