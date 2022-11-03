package getDashboard

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	DashboardRepository(user *model.User) (*model.User, *int64, *int64, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDashboard(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DashboardRepository(input *model.User) (*model.User, *int64, *int64, string) {

	var user model.User
	db := r.db.Model(&user)
	var count int64
	var activeCount int64
	errorCode := make(chan string, 1)

	db.Select("*").Count(&count)

	db.Select("*").Where("active = ?", true).Count(&activeCount)

	getUser := db.Select("*").Where("id = ?", input.ID).Find(&user)

	if getUser.RowsAffected < 1 {
		errorCode <- "DASHBOARD_NOT_FOUND_404"
	} else {
		errorCode <- "nil"
	}

	return &user, &count, &activeCount, <-errorCode

}
