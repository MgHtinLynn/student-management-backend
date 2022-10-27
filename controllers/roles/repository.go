package getRoles

import (
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	RoleRepository(Role *models.Role) (*[]models.Role, *int64, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRole(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RoleRepository(Role *models.Role) (*[]models.Role, *int64, string) {
	var roles []models.Role
	var count int64

	db := r.db.Model(&roles)

	getRoles := db.Select("*").Find(&roles).Count(&count)

	errorCode := make(chan string, 1)

	if getRoles.Error != nil {
		errorCode <- "RoleS_NOT_FOUND_404"
	} else {
		errorCode <- "nil"
	}

	return &roles, &count, <-errorCode

}
