package getExamResult

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type Service interface {
	GetExamResultService(input *InputGetExamResult) (*model.ExamResult, string)
}

type service struct {
	repository Repository
}

func NewServiceExamResult(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetExamResultService(input *InputGetExamResult) (*model.ExamResult, string) {

	var ExamResult = model.ExamResult{
		ID: input.ID,
	}
	getExamResult, errCreateStudent := s.repository.ExamResultRepository(&ExamResult)

	return getExamResult, errCreateStudent
}
