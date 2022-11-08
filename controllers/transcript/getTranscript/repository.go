package getTranscript

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	TranscriptRepository(Transcript *model.Transcript) (*model.Transcript, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryTranscript(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) TranscriptRepository(input *model.Transcript) (*model.Transcript, string) {

	var Transcript model.Transcript
	db := r.db.Model(&Transcript)
	errorCode := make(chan string, 1)

	getTranscript := db.Select("*").Preload("Student").Preload("Lecture").Where("id = ?", input.ID).Find(&Transcript)

	if getTranscript.RowsAffected < 1 {
		errorCode <- "RESULT_Transcript_NOT_FOUND_404"
		return &Transcript, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Transcript, <-errorCode

}
