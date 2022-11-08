package updateExamResult

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateExamResultRepository(input *model.ExamResult) (*model.ExamResult, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateExamResultRepository(input *model.ExamResult) (*model.ExamResult, string) {

	var examResult model.ExamResult
	var checkExamResult model.ExamResult
	db := r.db.Model(&examResult)
	checkDB := r.db.Model(&checkExamResult)
	errorCode := make(chan string, 1)

	checkEmailExist := checkDB.Select("*").Not("id = ?", input.ID).Where("subject_id = ? ", input.SubjectID).Where("student_id = ?", input.StudentId).Find(&checkExamResult)

	if checkEmailExist.RowsAffected > 0 {
		errorCode <- "UPDATE_ExamResult_CONFLICT_400"
		return &checkExamResult, <-errorCode
	}

	examResult.ID = input.ID
	examResult.Status = input.Status
	examResult.StudentId = input.StudentId
	examResult.SubjectID = input.SubjectID
	examResult.FilePath = input.FilePath

	updateExamResult := db.Select("*").Where("id = ?", input.ID).Updates(&examResult)

	if updateExamResult.Error != nil {
		errorCode <- "UPDATE_ExamResult_FAILED_403"
		return &examResult, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &examResult, <-errorCode
}
