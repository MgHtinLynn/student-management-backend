package updatePassword

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/MgHtinLynn/final-year-project-mcc/utils"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	UpdateUserRepository(input *model.User) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateUserRepository(input *model.User) (*model.User, string) {

	var user model.User
	db := r.db.Model(&user)
	errorCode := make(chan string, 1)

	user.ID = input.ID

	updateUser := db.Select("password", "email_verified_at", "active").Where("id = ?", input.ID).Updates(model.User{Password: utils.HashPassword(input.Password), Active: true, EmailVerifiedAt: &(&struct{ time.Time }{time.Now()}).Time})

	if updateUser.Error != nil {
		errorCode <- "UPDATE_USER_FAILED_403"
		return &user, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &user, <-errorCode
}
