package routes

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/certificate/createCertificate"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/certificate/getCertificate"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/certificate/getCertificates"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/certificate/updateCertificate"
	handleCreateCertificate "github.com/MgHtinLynn/final-year-project-mcc/handlers/certificate-handler/createCertificate"
	handlerGetCertificate "github.com/MgHtinLynn/final-year-project-mcc/handlers/certificate-handler/getCertificate"
	handlerGetCertificates "github.com/MgHtinLynn/final-year-project-mcc/handlers/certificate-handler/getCertificates"
	handleUpdateCertificate "github.com/MgHtinLynn/final-year-project-mcc/handlers/certificate-handler/updateCertificate"
	"github.com/MgHtinLynn/final-year-project-mcc/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCertificateRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All ExamResult Route
	*/

	certificatesRepository := getCertificates.NewRepositoryCertificates(db)
	service := getCertificates.NewServiceCertificates(certificatesRepository)
	certificatesHandler := handlerGetCertificates.NewHandlerGetCertificates(service)

	certificateRepository := getCertificate.NewRepositoryCertificate(db)
	certificateService := getCertificate.NewServiceCertificate(certificateRepository)
	certificateHandler := handlerGetCertificate.NewHandlerGetCertificate(certificateService)
	////////
	createRepository := createCertificate.NewRepositoryCreate(db)
	createService := createCertificate.NewServiceCreate(createRepository)
	createHandler := handleCreateCertificate.NewHandlerCreateCertificate(createService)
	////////
	updateRepository := updateCertificate.NewRepositoryUpdate(db)
	updateService := updateCertificate.NewServiceUpdate(updateRepository)
	updateHandler := handleUpdateCertificate.NewHandlerUpdateCertificate(updateService)
	//
	groupRoute := route.Group("/api/v1").Use(middlewares.Auth())
	groupRoute.GET("/certificates", certificatesHandler.GetCertificatesHandler)

	groupRoute.POST("/certificates", createHandler.CreateCertificateHandler)
	groupRoute.GET("/certificates/:id", certificateHandler.GetCertificateHandler)
	groupRoute.PUT("/certificates/:id", updateHandler.UpdateCertificateHandler)

}
