package getCertificates

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetCertificatesService(c *gin.Context) (*[]models.Certificate, *int64, string)
}

type service struct {
	repository Repository
}

func NewServiceCertificates(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCertificatesService(c *gin.Context) (*[]models.Certificate, *int64, string) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var Certificate models.Certificate

	getCertificateLists, count, errGetCertificates := s.repository.CertificatesRepository(&Certificate, &pagination, c)

	return getCertificateLists, count, errGetCertificates
}
