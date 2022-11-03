package getSubjects

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetSubjectsService(c *gin.Context) (*[]models.Subject, *int64, string)
}

type service struct {
	repository Repository
}

func NewServiceSubjects(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetSubjectsService(c *gin.Context) (*[]models.Subject, *int64, string) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var Subject models.Subject

	getSubjectLists, count, errGetSubjects := s.repository.SubjectsRepository(&Subject, &pagination, c)

	return getSubjectLists, count, errGetSubjects
}
