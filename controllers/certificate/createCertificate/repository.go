package createCertificate

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRepository(input *model.Certificate) (*model.Certificate, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateRepository(input *model.Certificate) (*model.Certificate, string) {

	var certificate model.Certificate
	db := r.db.Model(&certificate)
	errorCode := make(chan string, 1)

	checkTranscriptExist := db.Select("*").Where("lecture_id = ? ", input.LectureID).Where("student_id = ?", input.StudentId).Find(&certificate)

	if checkTranscriptExist.RowsAffected > 0 {
		errorCode <- "CREATE_Transcript_CONFLICT_409"
		return &certificate, <-errorCode
	}

	certificate.StudentId = input.StudentId
	certificate.LectureID = input.LectureID
	certificate.FilePath = input.FilePath

	addNewTranscript := db.Create(&certificate)

	if addNewTranscript.Error != nil {
		errorCode <- "CREATE_Transcript_FAILED_403"
		return &certificate, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &certificate, <-errorCode
}
