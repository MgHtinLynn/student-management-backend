package getCertificate

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetCertificateService(input *InputGetCertificate) (*model.Certificate, string)
}

type service struct {
	repository Repository
}

func (s *service) GetCertificatesService(c *gin.Context) (*[]model.Certificate, *int64, string) {
	//TODO implement me
	panic("implement me")
}

func NewServiceCertificate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCertificateService(input *InputGetCertificate) (*model.Certificate, string) {

	var Certificate = model.Certificate{
		ID: input.ID,
	}
	getCertificate, errCreateStudent := s.repository.CertificateRepository(&Certificate)

	return getCertificate, errCreateStudent
}
