package getUsers

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	UserRepository(user *models.User, p *models.Pagination) (*[]models.User, *int64, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UserRepository(user *models.User, p *models.Pagination) (*[]models.User, *int64, string) {
	var users []models.User
	var count int64

	offset := (p.Page - 1) * p.Limit
	queryBuilder := r.db.Limit(p.Limit).Offset(offset).Order(p.Sort)

	getUsers := queryBuilder.Model(&models.User{}).Where(user).Find(&users).Count(&count)

	errorCode := make(chan string, 1)

	if getUsers.Error != nil {
		errorCode <- "USERS_NOT_FOUND_404"
	} else {
		errorCode <- "nil"
	}

	return &users, &count, <-errorCode

}
