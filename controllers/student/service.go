package getStudents

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetStudentsService(c *gin.Context) (*[]models.User, *int64, string)
}

type service struct {
	repository Repository
}

func NewServiceStudents(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetStudentsService(c *gin.Context) (*[]models.User, *int64, string) {
	var User models.User
	getUserLists, count, errGetStudent := s.repository.StudentRepository(&User)
	return getUserLists, count, errGetStudent
}
