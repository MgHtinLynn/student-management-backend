package getSubject

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	SubjectRepository(Subject *model.Subject) (*model.Subject, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositorySubject(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) SubjectRepository(input *model.Subject) (*model.Subject, string) {

	var Subject model.Subject
	db := r.db.Model(&Subject)
	errorCode := make(chan string, 1)

	getSubject := db.Select("*").Preload("Teacher").Preload("Lecture").Where("id = ?", input.ID).Find(&Subject)

	if getSubject.RowsAffected < 1 {
		errorCode <- "RESULT_SUBJECT_NOT_FOUND_404"
		return &Subject, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Subject, <-errorCode

}
