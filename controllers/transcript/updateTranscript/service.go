package updateTranscript

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type Service interface {
	UpdateTranscriptService(input *InputUpdateTranscript) (*model.Transcript, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateTranscriptService(input *InputUpdateTranscript) (*model.Transcript, string) {

	Transcripts := model.Transcript{
		ID:        input.ID,
		StudentId: input.StudentId,
		LectureID: input.LectureID,
		FilePath:  input.FilePath,
	}

	resultUpdateTranscript, errCreateTranscript := s.repository.UpdateTranscriptRepository(&Transcripts)

	return resultUpdateTranscript, errCreateTranscript
}
