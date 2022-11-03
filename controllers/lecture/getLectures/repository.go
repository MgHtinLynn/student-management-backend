package getLectures

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	LectureRepository(Lecture *models.Lecture, p *models.Pagination, c *gin.Context) (*[]models.Lecture, *int64, string)
}

type repository struct {
	db *gorm.DB
}

type requestBdy struct {
	Page  string `form:"page"`
	Limit string `form:"limit"`
}

func NewRepositoryLecture(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LectureRepository(Lecture *models.Lecture, p *models.Pagination, c *gin.Context) (*[]models.Lecture, *int64, string) {
	var lectures []models.Lecture
	var count int64

	var requestBdy requestBdy
	errorCode := make(chan string, 1)

	_ = c.ShouldBind(&requestBdy)

	if (requestBdy.Page) == "" {
		db := r.db.Model(&lectures)
		getLectures := db.Select("*").Find(&lectures).Count(&count)
		if getLectures.Error != nil {
			errorCode <- "Lectures_NOT_FOUND_404"
		} else {
			errorCode <- "nil"
		}

		return &lectures, &count, <-errorCode
	} else {
		offset := (p.Page - 1) * p.Limit
		queryBuilder := r.db.Limit(p.Limit).Offset(offset).Order(p.Sort)
		getLectures := queryBuilder.Model(&models.Lecture{}).Joins("Tutor").Where(Lecture).Find(&lectures).Count(&count)
		if getLectures.Error != nil {
			errorCode <- "Lectures_NOT_FOUND_404"
		} else {
			errorCode <- "nil"
		}

		return &lectures, &count, <-errorCode
	}

}
