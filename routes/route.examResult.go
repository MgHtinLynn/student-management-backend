package routes

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/exam-result/createExamResult"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/exam-result/getExamResult"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/exam-result/getExamResults"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/exam-result/updateExamResult"
	getStudents "github.com/MgHtinLynn/final-year-project-mcc/controllers/student"
	handleCreateExamResult "github.com/MgHtinLynn/final-year-project-mcc/handlers/exam-result-handler/createExamResult"
	handlerGetExamResult "github.com/MgHtinLynn/final-year-project-mcc/handlers/exam-result-handler/getExamResult"
	handlerGetExamResults "github.com/MgHtinLynn/final-year-project-mcc/handlers/exam-result-handler/getExamResults"
	handlerGetStudents "github.com/MgHtinLynn/final-year-project-mcc/handlers/exam-result-handler/getStudents"
	handleUpdateExamResult "github.com/MgHtinLynn/final-year-project-mcc/handlers/exam-result-handler/updateExamResult"
	"github.com/MgHtinLynn/final-year-project-mcc/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitExamResultRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All ExamResult Route
	*/

	ExamResultsRepository := getExamResults.NewRepositoryExamResults(db)
	ExamResultsService := getExamResults.NewServiceExamResults(ExamResultsRepository)
	ExamResultsHandler := handlerGetExamResults.NewHandlerGetExamResults(ExamResultsService)

	StudentsRepository := getStudents.NewRepositoryStudents(db)
	StudentsService := getStudents.NewServiceStudents(StudentsRepository)
	StudentsHandler := handlerGetStudents.NewHandlerGetStudents(StudentsService)

	//
	ExamResultRepository := getExamResult.NewRepositoryExamResult(db)
	ExamResultService := getExamResult.NewServiceExamResult(ExamResultRepository)
	ExamResultHandler := handlerGetExamResult.NewHandlerGetExamResult(ExamResultService)
	//////
	createExamResultRepository := createExamResult.NewRepositoryCreate(db)
	createExamResultService := createExamResult.NewServiceCreate(createExamResultRepository)
	createExamResultHandler := handleCreateExamResult.NewHandlerCreateExamResult(createExamResultService)
	//////
	updateExamResultRepository := updateExamResult.NewRepositoryUpdate(db)
	updateExamResultService := updateExamResult.NewServiceUpdate(updateExamResultRepository)
	updateExamResultHandler := handleUpdateExamResult.NewHandlerUpdateExamResult(updateExamResultService)
	//
	groupRoute := route.Group("/api/v1").Use(middlewares.Auth())
	groupRoute.GET("/exam-results", ExamResultsHandler.GetExamResultsHandler)

	groupRoute.GET("/students", StudentsHandler.GetStudentsHandler)

	groupRoute.POST("/exam-results", createExamResultHandler.CreateExamResultHandler)
	groupRoute.GET("/exam-results/:id", ExamResultHandler.GetExamResultHandler)
	groupRoute.PUT("/exam-results/:id", updateExamResultHandler.UpdateExamResultHandler)

}
