package getTranscript

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTranscriptService(input *InputGetTranscript) (*model.Transcript, string)
}

type service struct {
	repository Repository
}

func (s *service) GetTranscriptsService(c *gin.Context) (*[]model.Transcript, *int64, string) {
	//TODO implement me
	panic("implement me")
}

func NewServiceTranscript(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetTranscriptService(input *InputGetTranscript) (*model.Transcript, string) {

	var Transcript = model.Transcript{
		ID: input.ID,
	}
	getTranscript, errCreateStudent := s.repository.TranscriptRepository(&Transcript)

	return getTranscript, errCreateStudent
}
