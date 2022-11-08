package createUser

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/MgHtinLynn/final-year-project-mcc/utils"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUserRepository(input *model.User) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateUserRepository(input *model.User) (*model.User, string) {

	var users model.User
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkEmailExist := db.Select("*").Not("id = ?", input.ID).Where("email = ?", input.Email).Find(&users)

	if checkEmailExist.RowsAffected > 0 {
		errorCode <- "CREATE_USER_EMAIL_CONFLICT_400"
		return &users, <-errorCode
	}

	checkUserExist := db.Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserExist.RowsAffected > 0 {
		errorCode <- "CREATE_USER_CONFLICT_409"
		return &users, <-errorCode
	}

	users.Name = input.Name
	users.Email = input.Email
	users.Phone = input.Phone
	users.Password = utils.HashPassword(input.Password)
	users.Active = input.Active
	users.Address = input.Address
	users.Role = input.Role
	users.ProfileUrl = input.ProfileUrl

	addNewStudent := db.Create(&users)

	if addNewStudent.Error != nil {
		errorCode <- "CREATE_STUDENT_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
