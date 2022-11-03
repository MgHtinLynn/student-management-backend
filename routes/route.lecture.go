package routes

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/lecture/createLecture"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/lecture/getLecture"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/lecture/getLectures"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/lecture/updateLecture"
	getTeachers "github.com/MgHtinLynn/final-year-project-mcc/controllers/teacher"
	handleCreateLecture "github.com/MgHtinLynn/final-year-project-mcc/handlers/lecture-handler/createLecture"
	handleResultLecture "github.com/MgHtinLynn/final-year-project-mcc/handlers/lecture-handler/getLecture"
	handleResultLectures "github.com/MgHtinLynn/final-year-project-mcc/handlers/lecture-handler/getLectures"
	handlerGetLectures "github.com/MgHtinLynn/final-year-project-mcc/handlers/lecture-handler/getTeachers"
	handleUpdateLecture "github.com/MgHtinLynn/final-year-project-mcc/handlers/lecture-handler/updateLecture"
	"github.com/MgHtinLynn/final-year-project-mcc/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitLectureRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Lecture Route
	*/

	LecturesRepository := getLectures.NewRepositoryLecture(db)
	LecturesService := getLectures.NewServiceLectures(LecturesRepository)
	LecturesHandler := handleResultLectures.NewHandlerGetLectures(LecturesService)

	TeachersRepository := getTeachers.NewRepositoryTeacher(db)
	TeachersService := getTeachers.NewServiceTeacher(TeachersRepository)
	TeachersHandler := handlerGetLectures.NewHandlerGetTeachers(TeachersService)

	LectureRepository := getLecture.NewRepositoryLecture(db)
	LectureService := getLecture.NewServiceLecture(LectureRepository)
	LectureHandler := handleResultLecture.NewHandlerGetLecture(LectureService)
	//
	createLectureRepository := createLecture.NewRepositoryCreate(db)
	createLectureService := createLecture.NewServiceCreate(createLectureRepository)
	createLectureHandler := handleCreateLecture.NewHandlerCreateLecture(createLectureService)
	//
	updateLectureRepository := updateLecture.NewRepositoryUpdate(db)
	updateLectureService := updateLecture.NewServiceUpdate(updateLectureRepository)
	updateLectureHandler := handleUpdateLecture.NewHandlerUpdateLecture(updateLectureService)

	groupRoute := route.Group("/api/v1").Use(middlewares.Auth())
	groupRoute.GET("/lectures", LecturesHandler.GetLecturesHandler)

	groupRoute.GET("/teachers", TeachersHandler.GetTeachersHandler)

	groupRoute.POST("/lectures", createLectureHandler.CreateLectureHandler)
	groupRoute.GET("/lectures/:id", LectureHandler.GetLectureHandler)
	groupRoute.PUT("/lectures/:id", updateLectureHandler.UpdateLectureHandler)

}
