package getExamResults

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	ExamResultsRepository(ExamResult *models.ExamResult, p *models.Pagination, c *gin.Context) (*[]models.ExamResult, *int64, string)
}

type repository struct {
	db *gorm.DB
}

type requestBdy struct {
	Page  string `form:"page" binding:"required"`
	Limit string `form:"limit" binding:"required"`
}

func NewRepositoryExamResults(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ExamResultsRepository(ExamResult *models.ExamResult, p *models.Pagination, c *gin.Context) (*[]models.ExamResult, *int64, string) {
	var examResults []models.ExamResult
	var count int64

	var requestBdy requestBdy
	errorCode := make(chan string, 1)

	_ = c.ShouldBind(&requestBdy)

	if (requestBdy.Page) == "" {
		db := r.db.Model(&examResults)
		getExamResults := db.Select("*").Find(&examResults).Count(&count)
		if getExamResults.Error != nil {
			errorCode <- "ExamResults_NOT_FOUND_404"
		} else {
			errorCode <- "nil"
		}

		return &examResults, &count, <-errorCode
	} else {
		offset := (p.Page - 1) * p.Limit
		queryBuilder := r.db.Limit(p.Limit).Offset(offset).Order(p.Sort)
		getExamResults := queryBuilder.Model(&models.ExamResult{}).Preload("Student").Preload("Subject").Where(ExamResult).Find(&examResults).Count(&count)
		if getExamResults.Error != nil {
			errorCode <- "ExamResults_NOT_FOUND_404"
		} else {
			errorCode <- "nil"
		}

		return &examResults, &count, <-errorCode
	}

}
