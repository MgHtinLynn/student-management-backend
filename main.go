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
	"net/http"
)

func main() {
	/**
	@description Setup Server
	*/
	router := SetupRouter()
	/**
	@description Run Server
	*/
	log.Fatal(router.Run(":" + utils.LoadConfig("GO_PORT")))
}

func SetupRouter() *gin.Engine {
	/**
	@description Setup Database connection
	*/
	db := config.Connection()

	/**
	@description Setup Mode Application
	*/
	if utils.LoadConfig("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	/**
	@description Init Router
	*/
	r := gin.Default()

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
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "It's worked",
		})
	})

	route.InitAuthRoutes(db, r)
	route.InitUserRoutes(db, r)
	route.InitRoleRoutes(db, r)

	return r
}
