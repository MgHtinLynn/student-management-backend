package getRoles

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetRolesService(c *gin.Context) (*[]models.Role, *int64, string)
}

type service struct {
	repository Repository
}

func NewServiceRoles(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetRolesService(c *gin.Context) (*[]models.Role, *int64, string) {
	var Role models.Role
	getRoleLists, count, errGetRoles := s.repository.RoleRepository(&Role)
	return getRoleLists, count, errGetRoles
}
