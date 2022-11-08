package updateExamResult

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type Service interface {
	UpdateExamResultService(input *InputUpdateExamResult) (*model.ExamResult, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateExamResultService(input *InputUpdateExamResult) (*model.ExamResult, string) {

	ExamResults := model.ExamResult{
		ID:        input.ID,
		Status:    input.Status,
		StudentId: input.StudentId,
		SubjectID: input.SubjectID,
		FilePath:  input.FilePath,
	}

	resultUpdateExamResult, errCreateExamResult := s.repository.UpdateExamResultRepository(&ExamResults)

	return resultUpdateExamResult, errCreateExamResult
}
