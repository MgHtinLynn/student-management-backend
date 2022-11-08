package getExamResults

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetExamResultsService(c *gin.Context) (*[]models.ExamResult, *int64, string)
}

type service struct {
	repository Repository
}

func (s *service) GetCertificatesService(c *gin.Context) (*[]models.Certificate, *int64, string) {
	//TODO implement me
	panic("implement me")
}

func NewServiceExamResults(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetExamResultsService(c *gin.Context) (*[]models.ExamResult, *int64, string) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var ExamResult models.ExamResult

	getExamResultLists, count, errGetExamResults := s.repository.ExamResultsRepository(&ExamResult, &pagination, c)

	return getExamResultLists, count, errGetExamResults
}
