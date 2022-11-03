package getUser

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	UserRepository(user *model.User) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UserRepository(input *model.User) (*model.User, string) {

	var user model.User
	db := r.db.Model(&user)
	errorCode := make(chan string, 1)

	getUser := db.Select("*").Where("id = ?", input.ID).Find(&user)

	if getUser.RowsAffected < 1 {
		errorCode <- "RESULT_STUDENT_NOT_FOUND_404"
		return &user, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &user, <-errorCode

}
