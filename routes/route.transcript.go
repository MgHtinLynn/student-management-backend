package routes

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/transcript/createTranscript"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/transcript/getTranscript"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/transcript/getTranscripts"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/transcript/updateTranscript"
	handleCreateTranscript "github.com/MgHtinLynn/final-year-project-mcc/handlers/transcript-handler/createTranscript"
	handlerGetTranscript "github.com/MgHtinLynn/final-year-project-mcc/handlers/transcript-handler/getTranscript"
	handlerGetTranscripts "github.com/MgHtinLynn/final-year-project-mcc/handlers/transcript-handler/getTranscripts"
	handleUpdateTranscript "github.com/MgHtinLynn/final-year-project-mcc/handlers/transcript-handler/updateTranscript"
	"github.com/MgHtinLynn/final-year-project-mcc/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitTranscriptRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All ExamResult Route
	*/

	transcriptsRepository := getTranscripts.NewRepositoryTranscripts(db)
	service := getTranscripts.NewServiceTranscripts(transcriptsRepository)
	transcriptsHandler := handlerGetTranscripts.NewHandlerGetTranscripts(service)
	//
	////
	transcriptRepository := getTranscript.NewRepositoryTranscript(db)
	transcriptService := getTranscript.NewServiceTranscript(transcriptRepository)
	transcriptHandler := handlerGetTranscript.NewHandlerGetTranscript(transcriptService)
	////////
	createRepository := createTranscript.NewRepositoryCreate(db)
	createService := createTranscript.NewServiceCreate(createRepository)
	createHandler := handleCreateTranscript.NewHandlerCreateTranscript(createService)
	////////
	updateRepository := updateTranscript.NewRepositoryUpdate(db)
	updateService := updateTranscript.NewServiceUpdate(updateRepository)
	updateHandler := handleUpdateTranscript.NewHandlerUpdateTranscript(updateService)
	//
	groupRoute := route.Group("/api/v1").Use(middlewares.Auth())
	groupRoute.GET("/transcripts", transcriptsHandler.GetTranscriptsHandler)

	groupRoute.POST("/transcripts", createHandler.CreateTranscriptHandler)
	groupRoute.GET("/transcripts/:id", transcriptHandler.GetTranscriptHandler)
	groupRoute.PUT("/transcripts/:id", updateHandler.UpdateTranscriptHandler)

}
