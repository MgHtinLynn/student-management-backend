package getUsers

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetUsersService(c *gin.Context) (*[]models.User, *int64, string)
}

type service struct {
	repository Repository
}

func NewServiceUsers(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetUsersService(c *gin.Context) (*[]models.User, *int64, string) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var user models.User
	getUserLists, count, errGetUsers := s.repository.UserRepository(&user, &pagination)
	return getUserLists, count, errGetUsers
}
