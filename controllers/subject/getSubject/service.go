package getSubject

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type Service interface {
	GetSubjectService(input *InputGetSubject) (*model.Subject, string)
}

type service struct {
	repository Repository
}

func NewServiceSubject(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetSubjectService(input *InputGetSubject) (*model.Subject, string) {

	var Subject = model.Subject{
		ID: input.ID,
	}
	getSubject, errCreateStudent := s.repository.SubjectRepository(&Subject)

	return getSubject, errCreateStudent
}
