package routes

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/subject/createSubject"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/subject/getSubject"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/subject/getSubjects"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/subject/updateSubject"
	handleCreateSubject "github.com/MgHtinLynn/final-year-project-mcc/handlers/subject-handler/createSubject"
	handlerGetSubject "github.com/MgHtinLynn/final-year-project-mcc/handlers/subject-handler/getSubject"
	handlerGetSubjects "github.com/MgHtinLynn/final-year-project-mcc/handlers/subject-handler/getSubjects"
	handleUpdateSubject "github.com/MgHtinLynn/final-year-project-mcc/handlers/subject-handler/updateSubject"
	"github.com/MgHtinLynn/final-year-project-mcc/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitSubjectRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Subject Route
	*/

	SubjectsRepository := getSubjects.NewRepositorySubject(db)
	SubjectsService := getSubjects.NewServiceSubjects(SubjectsRepository)
	SubjectsHandler := handlerGetSubjects.NewHandlerGetSubjects(SubjectsService)

	//
	SubjectRepository := getSubject.NewRepositorySubject(db)
	SubjectService := getSubject.NewServiceSubject(SubjectRepository)
	SubjectHandler := handlerGetSubject.NewHandlerGetSubject(SubjectService)
	////
	createSubjectRepository := createSubject.NewRepositoryCreate(db)
	createSubjectService := createSubject.NewServiceCreate(createSubjectRepository)
	createSubjectHandler := handleCreateSubject.NewHandlerCreateSubject(createSubjectService)
	////
	updateSubjectRepository := updateSubject.NewRepositoryUpdate(db)
	updateSubjectService := updateSubject.NewServiceUpdate(updateSubjectRepository)
	updateSubjectHandler := handleUpdateSubject.NewHandlerUpdateSubject(updateSubjectService)
	//
	groupRoute := route.Group("/api/v1").Use(middlewares.Auth())
	groupRoute.GET("/subjects", SubjectsHandler.GetSubjectsHandler)

	groupRoute.POST("/subjects", createSubjectHandler.CreateSubjectHandler)
	groupRoute.GET("/subjects/:id", SubjectHandler.GetSubjectHandler)
	groupRoute.PUT("/subjects/:id", updateSubjectHandler.UpdateSubjectHandler)

}
