package getTranscripts

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTranscriptsService(c *gin.Context) (*[]models.Transcript, *int64, string)
}

type service struct {
	repository Repository
}

func NewServiceTranscripts(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetTranscriptsService(c *gin.Context) (*[]models.Transcript, *int64, string) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var Transcript models.Transcript

	getTranscriptLists, count, errGetTranscripts := s.repository.TranscriptsRepository(&Transcript, &pagination, c)

	return getTranscriptLists, count, errGetTranscripts
}
