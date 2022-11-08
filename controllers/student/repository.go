package getStudents

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	StudentRepository(user *models.User) (*[]models.User, *int64, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryStudents(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) StudentRepository(user *models.User) (*[]models.User, *int64, string) {
	var users []models.User
	var count int64

	db := r.db.Model(&users)

	getStudents := db.Select("*").Where("role = ? ", "student").Where("active = ? ", true).Find(&users).Count(&count)

	errorCode := make(chan string, 1)

	if getStudents.Error != nil {
		errorCode <- "StudentS_NOT_FOUND_404"
	} else {
		errorCode <- "nil"
	}

	return &users, &count, <-errorCode

}
