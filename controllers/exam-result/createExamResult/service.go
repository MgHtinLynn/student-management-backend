package createExamResult

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type Service interface {
	CreateExamResultService(input *InputCreateExamResult) (*model.ExamResult, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateExamResultService(input *InputCreateExamResult) (*model.ExamResult, string) {

	ExamResults := model.ExamResult{
		Status:    input.Status,
		StudentId: input.StudentId,
		SubjectID: input.SubjectID,
		FilePath:  input.FilePath,
	}

	resultCreateUser, errCreateUser := s.repository.CreateExamResultRepository(&ExamResults)

	return resultCreateUser, errCreateUser
}
