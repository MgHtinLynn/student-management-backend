package loginAuth

import (
	"fmt"
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"gorm.io/gorm"
)

type Repository interface {
	LoginRepository(input *models.User) (*models.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LoginRepository(input *models.User) (*models.User, string) {
	var users models.User
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email
	users.Password = input.Password
	fmt.Println("email\t\t", input.Email)

	checkUserAccount := db.Select("*").Where("email = ?", input.Email).Find(&users)
	fmt.Println(checkUserAccount)
	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	ComparePassword := util.ComparePassword(users.Password, input.Password)

	if ComparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode

}
