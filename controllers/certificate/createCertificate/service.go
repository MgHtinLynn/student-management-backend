package createCertificate

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type Service interface {
	CreateService(input *InputCreateCertificate) (*model.Certificate, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateService(input *InputCreateCertificate) (*model.Certificate, string) {

	Certificates := model.Certificate{
		StudentId: input.StudentId,
		LectureID: input.LectureID,
		FilePath:  input.FilePath,
	}

	resultCreateUser, errCreateUser := s.repository.CreateRepository(&Certificates)

	return resultCreateUser, errCreateUser
}
