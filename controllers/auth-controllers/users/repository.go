package getUsers

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	UserRepository(user *models.User, p *models.Pagination, c *gin.Context) (*[]models.User, *int64, string)
}

type repository struct {
	db *gorm.DB
}

type requestBdy struct {
	Role  string `form:"role"`
	Email string `form:"email"`
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UserRepository(user *models.User, p *models.Pagination, c *gin.Context) (*[]models.User, *int64, string) {
	var users []models.User
	var count int64

	var requestBdy requestBdy
	errorCode := make(chan string, 1)

	_ = c.ShouldBind(&requestBdy)

	if (requestBdy.Email) != "" {
		db := r.db.Model(&users)
		getUser := db.Select("*").Where("email = ?", requestBdy.Email).Find(&users).Count(&count)

		if getUser.Error != nil {
			errorCode <- "USERS_NOT_FOUND_404"
		} else {
			errorCode <- "nil"
		}

		return &users, &count, <-errorCode
	}

	offset := (p.Page - 1) * p.Limit
	queryBuilder := r.db.Limit(p.Limit).Offset(offset).Order(p.Sort)

	getUsers := queryBuilder.Model(&models.User{}).Where(user).Find(&users).Count(&count)

	if getUsers.Error != nil {
		errorCode <- "USERS_NOT_FOUND_404"
	} else {
		errorCode <- "nil"
	}

	return &users, &count, <-errorCode

}
