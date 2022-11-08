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
	var checkUser model.User
	db := r.db.Model(&user)
	checkDB := r.db.Model(&user)
	errorCode := make(chan string, 1)

	checkEmailExist := checkDB.Select("*").Not("id = ?", input.ID).Where("email = ?", input.Email).Find(&checkUser)

	if checkEmailExist.RowsAffected > 0 {
		errorCode <- "UPDATE_USER_EMAIL_CONFLICT_400"
		return &user, <-errorCode
	}

	user.ID = input.ID
	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone
	user.Active = input.Active
	user.Address = input.Address
	user.Role = input.Role
	user.ProfileUrl = input.ProfileUrl

	fmt.Println("user", user)

	updateUser := db.Select("*").Where("id = ?", input.ID).Updates(&user)

	if updateUser.Error != nil {
		errorCode <- "UPDATE_USER_FAILED_403"
		return &user, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &user, <-errorCode
}
