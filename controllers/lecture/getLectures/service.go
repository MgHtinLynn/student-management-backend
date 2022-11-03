package getLectures

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetLecturesService(c *gin.Context) (*[]models.Lecture, *int64, string)
}

type service struct {
	repository Repository
}

func NewServiceLectures(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetLecturesService(c *gin.Context) (*[]models.Lecture, *int64, string) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var Lecture models.Lecture
	getLectureLists, count, errGetLectures := s.repository.LectureRepository(&Lecture, &pagination, c)
	return getLectureLists, count, errGetLectures
}
