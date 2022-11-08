package updateCertificate

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateCertificateRepository(input *model.Certificate) (*model.Certificate, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateCertificateRepository(input *model.Certificate) (*model.Certificate, string) {

	var examResult model.Certificate
	var checkCertificate model.Certificate
	db := r.db.Model(&examResult)
	checkDB := r.db.Model(&checkCertificate)
	errorCode := make(chan string, 1)

	checkEmailExist := checkDB.Select("*").Not("id = ?", input.ID).Where("lecture_id = ? ", input.LectureID).Where("student_id = ?", input.StudentId).Find(&checkCertificate)

	if checkEmailExist.RowsAffected > 0 {
		errorCode <- "UPDATE_Certificate_CONFLICT_400"
		return &checkCertificate, <-errorCode
	}

	examResult.ID = input.ID
	examResult.StudentId = input.StudentId
	examResult.LectureID = input.LectureID
	examResult.FilePath = input.FilePath

	updateCertificate := db.Select("*").Where("id = ?", input.ID).Updates(&examResult)

	if updateCertificate.Error != nil {
		errorCode <- "UPDATE_Certificate_FAILED_403"
		return &examResult, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &examResult, <-errorCode
}
