package updateTranscript

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateTranscriptRepository(input *model.Transcript) (*model.Transcript, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateTranscriptRepository(input *model.Transcript) (*model.Transcript, string) {

	var examResult model.Transcript
	var checkTranscript model.Transcript
	db := r.db.Model(&examResult)
	checkDB := r.db.Model(&checkTranscript)
	errorCode := make(chan string, 1)

	checkEmailExist := checkDB.Select("*").Not("id = ?", input.ID).Where("lecture_id = ? ", input.LectureID).Where("student_id = ?", input.StudentId).Find(&checkTranscript)

	if checkEmailExist.RowsAffected > 0 {
		errorCode <- "UPDATE_Transcript_CONFLICT_400"
		return &checkTranscript, <-errorCode
	}

	examResult.ID = input.ID
	examResult.StudentId = input.StudentId
	examResult.LectureID = input.LectureID
	examResult.FilePath = input.FilePath

	updateTranscript := db.Select("*").Where("id = ?", input.ID).Updates(&examResult)

	if updateTranscript.Error != nil {
		errorCode <- "UPDATE_Transcript_FAILED_403"
		return &examResult, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &examResult, <-errorCode
}
