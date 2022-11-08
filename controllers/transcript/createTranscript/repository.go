package createTranscript

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRepository(input *model.Transcript) (*model.Transcript, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateRepository(input *model.Transcript) (*model.Transcript, string) {

	var Transcripts model.Transcript
	db := r.db.Model(&Transcripts)
	errorCode := make(chan string, 1)

	checkTranscriptExist := db.Select("*").Where("lecture_id = ? ", input.LectureID).Where("student_id = ?", input.StudentId).Find(&Transcripts)

	if checkTranscriptExist.RowsAffected > 0 {
		errorCode <- "CREATE_Transcript_CONFLICT_409"
		return &Transcripts, <-errorCode
	}

	Transcripts.StudentId = input.StudentId
	Transcripts.LectureID = input.LectureID
	Transcripts.FilePath = input.FilePath

	addNewTranscript := db.Create(&Transcripts)

	if addNewTranscript.Error != nil {
		errorCode <- "CREATE_Transcript_FAILED_403"
		return &Transcripts, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Transcripts, <-errorCode
}
