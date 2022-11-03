package updateUser

import (
	"fmt"
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
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

	checkEmailExist := db.Select("*").Not("id = ?", input.ID).Where("email = ?", input.Email).Find(&user)

	if checkEmailExist.RowsAffected > 0 {
		errorCode <- "UPDATE_USER_EMAIL_CONFLICT_400"
		return &user, <-errorCode
	}

	user.ID = input.ID

	checkUserId := db.Select("*").Where("id = ?", input.ID).Find(&user)

	if checkUserId.RowsAffected < 1 {
		errorCode <- "UPDATE_USER_NOT_FOUND_404"
		return &user, <-errorCode
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone
	user.Active = input.Active
	user.Address = input.Address
	user.Role = input.Role
	user.ProfileUrl = input.ProfileUrl

	fmt.Println("user", user)

	updateUser := db.Select("*").Where("id = ?", input.ID).Updates(user)

	fmt.Println("updateUser", updateUser)

	if updateUser.Error != nil {
		errorCode <- "UPDATE_USER_FAILED_403"
		return &user, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &user, <-errorCode
}
