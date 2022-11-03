package getTeachers

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTeachersService(c *gin.Context) (*[]models.User, *int64, string)
}

type service struct {
	repository Repository
}

func NewServiceTeacher(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetTeachersService(c *gin.Context) (*[]models.User, *int64, string) {
	var User models.User
	getUserLists, count, errGetTeacher := s.repository.TeacherRepository(&User)
	return getUserLists, count, errGetTeacher
}
