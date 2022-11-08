package getTranscripts

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	TranscriptsRepository(Transcript *models.Transcript, p *models.Pagination, c *gin.Context) (*[]models.Transcript, *int64, string)
}

type repository struct {
	db *gorm.DB
}

type requestBdy struct {
	Page  string `form:"page" binding:"required"`
	Limit string `form:"limit" binding:"required"`
}

func NewRepositoryTranscripts(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) TranscriptsRepository(Transcript *models.Transcript, p *models.Pagination, c *gin.Context) (*[]models.Transcript, *int64, string) {
	var Transcripts []models.Transcript
	var count int64

	var requestBdy requestBdy
	errorCode := make(chan string, 1)

	_ = c.ShouldBind(&requestBdy)

	if (requestBdy.Page) == "" {
		db := r.db.Model(&Transcripts)
		getTranscripts := db.Select("*").Find(&Transcripts).Count(&count)
		if getTranscripts.Error != nil {
			errorCode <- "Transcripts_NOT_FOUND_404"
		} else {
			errorCode <- "nil"
		}

		return &Transcripts, &count, <-errorCode
	} else {
		offset := (p.Page - 1) * p.Limit
		queryBuilder := r.db.Limit(p.Limit).Offset(offset).Order(p.Sort)
		getTranscripts := queryBuilder.Model(&models.Transcript{}).Preload("Student").Preload("Lecture").Where(Transcript).Find(&Transcripts).Count(&count)
		if getTranscripts.Error != nil {
			errorCode <- "Transcripts_NOT_FOUND_404"
		} else {
			errorCode <- "nil"
		}

		return &Transcripts, &count, <-errorCode
	}

}
