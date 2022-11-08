package getAttendances

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	AttendanceRepository(Attendance *models.Attendance, p *models.Pagination) (*[]models.Attendance, *int64, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryAttendance(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) AttendanceRepository(Attendance *models.Attendance, p *models.Pagination) (*[]models.Attendance, *int64, string) {
	var attendances []models.Attendance
	var count int64

	offset := (p.Page - 1) * p.Limit
	queryBuilder := r.db.Limit(p.Limit).Offset(offset).Order(p.Sort)
	getAttendances := queryBuilder.Model(&models.Attendance{}).Preload("Student").Preload("Subject").Where(Attendance).Find(&attendances).Count(&count)

	errorCode := make(chan string, 1)

	if getAttendances.Error != nil {
		errorCode <- "Attendances_NOT_FOUND_404"
	} else {
		errorCode <- "nil"
	}

	return &attendances, &count, <-errorCode

}
