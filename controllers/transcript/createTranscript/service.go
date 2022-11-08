package createTranscript

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type Service interface {
	CreateService(input *InputCreateTranscript) (*model.Transcript, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateService(input *InputCreateTranscript) (*model.Transcript, string) {

	Transcripts := model.Transcript{
		StudentId: input.StudentId,
		LectureID: input.LectureID,
		FilePath:  input.FilePath,
	}

	resultCreateUser, errCreateUser := s.repository.CreateRepository(&Transcripts)

	return resultCreateUser, errCreateUser
}
