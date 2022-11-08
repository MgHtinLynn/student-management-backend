package getAttendances

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAttendancesService(c *gin.Context) (*[]models.Attendance, *int64, string)
}

type service struct {
	repository Repository
}

func NewServiceAttendances(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAttendancesService(c *gin.Context) (*[]models.Attendance, *int64, string) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var Attendance models.Attendance
	getAttendanceLists, count, errGetAttendances := s.repository.AttendanceRepository(&Attendance, &pagination)
	return getAttendanceLists, count, errGetAttendances
}
