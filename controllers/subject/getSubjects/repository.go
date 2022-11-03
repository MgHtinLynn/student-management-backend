package getSubjects

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	SubjectsRepository(Subject *models.Subject, p *models.Pagination, c *gin.Context) (*[]models.Subject, *int64, string)
}

type repository struct {
	db *gorm.DB
}

type requestBdy struct {
	Page  string `form:"page" binding:"required"`
	Limit string `form:"limit" binding:"required"`
}

func NewRepositorySubject(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) SubjectsRepository(Subject *models.Subject, p *models.Pagination, c *gin.Context) (*[]models.Subject, *int64, string) {
	var subjects []models.Subject
	var count int64

	var requestBdy requestBdy
	errorCode := make(chan string, 1)

	_ = c.ShouldBind(&requestBdy)

	if (requestBdy.Page) == "" {
		db := r.db.Model(&subjects)
		getSubjects := db.Select("*").Find(&subjects).Count(&count)
		if getSubjects.Error != nil {
			errorCode <- "Subjects_NOT_FOUND_404"
		} else {
			errorCode <- "nil"
		}

		return &subjects, &count, <-errorCode
	} else {
		offset := (p.Page - 1) * p.Limit
		queryBuilder := r.db.Limit(p.Limit).Offset(offset).Order(p.Sort)
		getSubjects := queryBuilder.Model(&models.Subject{}).Preload("Teacher").Preload("Lecture").Where(Subject).Find(&subjects).Count(&count)
		if getSubjects.Error != nil {
			errorCode <- "Subjects_NOT_FOUND_404"
		} else {
			errorCode <- "nil"
		}

		return &subjects, &count, <-errorCode
	}

}
