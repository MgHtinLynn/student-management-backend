package getTeachers

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	TeacherRepository(user *models.User) (*[]models.User, *int64, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryTeacher(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) TeacherRepository(user *models.User) (*[]models.User, *int64, string) {
	var users []models.User
	var count int64

	db := r.db.Model(&users)

	getTeachers := db.Select("*").Where("role = ? ", "teacher").Where("active = ? ", true).Find(&users).Count(&count)

	errorCode := make(chan string, 1)

	if getTeachers.Error != nil {
		errorCode <- "TEACHERS_NOT_FOUND_404"
	} else {
		errorCode <- "nil"
	}

	return &users, &count, <-errorCode

}
