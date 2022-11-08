package getCertificates

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	CertificatesRepository(Certificate *models.Certificate, p *models.Pagination, c *gin.Context) (*[]models.Certificate, *int64, string)
}

type repository struct {
	db *gorm.DB
}

type requestBdy struct {
	Page  string `form:"page" binding:"required"`
	Limit string `form:"limit" binding:"required"`
}

func NewRepositoryCertificates(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CertificatesRepository(Certificate *models.Certificate, p *models.Pagination, c *gin.Context) (*[]models.Certificate, *int64, string) {
	var Certificates []models.Certificate
	var count int64

	var requestBdy requestBdy
	errorCode := make(chan string, 1)

	_ = c.ShouldBind(&requestBdy)

	if (requestBdy.Page) == "" {
		db := r.db.Model(&Certificates)
		getCertificates := db.Select("*").Find(&Certificates).Count(&count)
		if getCertificates.Error != nil {
			errorCode <- "Certificates_NOT_FOUND_404"
		} else {
			errorCode <- "nil"
		}

		return &Certificates, &count, <-errorCode
	} else {
		offset := (p.Page - 1) * p.Limit
		queryBuilder := r.db.Limit(p.Limit).Offset(offset).Order(p.Sort)
		getCertificates := queryBuilder.Model(&models.Certificate{}).Preload("Student").Preload("Lecture").Where(Certificate).Find(&Certificates).Count(&count)
		if getCertificates.Error != nil {
			errorCode <- "Certificates_NOT_FOUND_404"
		} else {
			errorCode <- "nil"
		}

		return &Certificates, &count, <-errorCode
	}

}
