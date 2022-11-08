package getExamResult

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	ExamResultRepository(ExamResult *model.ExamResult) (*model.ExamResult, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryExamResult(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ExamResultRepository(input *model.ExamResult) (*model.ExamResult, string) {

	var ExamResult model.ExamResult
	db := r.db.Model(&ExamResult)
	errorCode := make(chan string, 1)

	getExamResult := db.Select("*").Preload("Student").Preload("Subject").Where("id = ?", input.ID).Find(&ExamResult)

	if getExamResult.RowsAffected < 1 {
		errorCode <- "RESULT_EXAM_RESULT_NOT_FOUND_404"
		return &ExamResult, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &ExamResult, <-errorCode

}
