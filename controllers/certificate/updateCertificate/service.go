package updateCertificate

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type Service interface {
	UpdateCertificateService(input *InputUpdateCertificate) (*model.Certificate, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateCertificateService(input *InputUpdateCertificate) (*model.Certificate, string) {

	Certificates := model.Certificate{
		ID:        input.ID,
		StudentId: input.StudentId,
		LectureID: input.LectureID,
		FilePath:  input.FilePath,
	}

	resultUpdateCertificate, errCreateCertificate := s.repository.UpdateCertificateRepository(&Certificates)

	return resultUpdateCertificate, errCreateCertificate
}
