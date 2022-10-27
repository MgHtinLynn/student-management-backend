package main

import (
	config "github.com/MgHtinLynn/final-year-project-mcc/configs"
	route "github.com/MgHtinLynn/final-year-project-mcc/routes"
	"github.com/MgHtinLynn/final-year-project-mcc/utils"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	/**
	@description Setup Server
	*/
	router := SetupRouter()
	/**
	@description Run Server
	*/
	log.Fatal(router.Run("localhost:" + utils.LoadConfig("GO_PORT")))
}

func SetupRouter() *gin.Engine {
	/**
	@description Setup Database connection
	*/
	db := config.Connection()
	/**
	@description Init Router
	*/
	r := gin.Default()

	/**
	@description Setup Mode Application
	*/
	//if utils.LoadConfig("GO_ENV") != "production" && utils.LoadConfig("GO_ENV") != "test" {
	gin.SetMode(gin.DebugMode)
	//} else {
	// gin.SetMode(gin.ReleaseMode)
	//}

	/**
	@description Setup Middleware
	*/

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	r.Use(helmet.Default())
	r.Use(gzip.Gzip(gzip.BestCompression))

	/**
	@description Init All Route
	*/
	route.InitAuthRoutes(db, r)
	route.InitUserRoutes(db, r)
	route.InitRoleRoutes(db, r)

	return r
}
