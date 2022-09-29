package routes

import (
	loginAuth "github.com/MgHtinLynn/final-year-project-mcc/controllers/auth-controllers/login"
	handlerLogin "github.com/MgHtinLynn/final-year-project-mcc/handlers/auth-handler/login"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	/**
	@description All Handler Auth
	*/

	LoginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(LoginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/login", loginHandler.LoginHandler)
}
