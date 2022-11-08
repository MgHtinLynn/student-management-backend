package createExamResult

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateExamResultRepository(input *model.ExamResult) (*model.ExamResult, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateExamResultRepository(input *model.ExamResult) (*model.ExamResult, string) {

	var ExamResults model.ExamResult
	db := r.db.Model(&ExamResults)
	errorCode := make(chan string, 1)

	checkExamResultExist := db.Select("*").Where("subject_id = ? ", input.SubjectID).Where("student_id = ?", input.StudentId).Find(&ExamResults)

	if checkExamResultExist.RowsAffected > 0 {
		errorCode <- "CREATE_ExamResult_CONFLICT_409"
		return &ExamResults, <-errorCode
	}

	ExamResults.Status = input.Status
	ExamResults.StudentId = input.StudentId
	ExamResults.SubjectID = input.SubjectID
	ExamResults.FilePath = input.FilePath

	addNewExamResult := db.Create(&ExamResults)

	if addNewExamResult.Error != nil {
		errorCode <- "CREATE_ExamResult_FAILED_403"
		return &ExamResults, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &ExamResults, <-errorCode
}
