package getCertificate

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	CertificateRepository(Certificate *model.Certificate) (*model.Certificate, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCertificate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CertificateRepository(input *model.Certificate) (*model.Certificate, string) {

	var Certificate model.Certificate
	db := r.db.Model(&Certificate)
	errorCode := make(chan string, 1)

	getCertificate := db.Select("*").Preload("Student").Preload("Lecture").Where("id = ?", input.ID).Find(&Certificate)

	if getCertificate.RowsAffected < 1 {
		errorCode <- "RESULT_Certificate_NOT_FOUND_404"
		return &Certificate, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Certificate, <-errorCode

}
