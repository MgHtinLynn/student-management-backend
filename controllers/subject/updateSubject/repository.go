package updateSubject

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateSubjectRepository(input *model.Subject) (*model.Subject, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateSubjectRepository(input *model.Subject) (*model.Subject, string) {

	var Subject model.Subject
	db := r.db.Model(&Subject)
	errorCode := make(chan string, 1)

	checkEmailExist := db.Select("*").Not("id = ?", input.ID).Where("name = ?", input.Name).Find(&Subject)

	if checkEmailExist.RowsAffected > 0 {
		errorCode <- "UPDATE_Subject_EMAIL_CONFLICT_400"
		return &Subject, <-errorCode
	}

	Subject.ID = input.ID

	checkSubjectId := db.Select("*").Where("id = ?", input.ID).Find(&Subject)

	if checkSubjectId.RowsAffected < 1 {
		errorCode <- "UPDATE_Subject_NOT_FOUND_404"
		return &Subject, <-errorCode
	}

	Subject.Name = input.Name
	Subject.TeacherId = input.TeacherId
	Subject.LectureId = input.LectureId

	updateSubject := db.Select("*").Where("id = ?", input.ID).Updates(Subject)

	if updateSubject.Error != nil {
		errorCode <- "UPDATE_Subject_FAILED_403"
		return &Subject, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Subject, <-errorCode
}
