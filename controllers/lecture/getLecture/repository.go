package getLecture

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	LectureRepository(Lecture *model.Lecture) (*model.Lecture, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryLecture(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LectureRepository(input *model.Lecture) (*model.Lecture, string) {

	var Lecture model.Lecture
	db := r.db.Model(&Lecture)
	errorCode := make(chan string, 1)

	getLecture := db.Select("*").Where("id = ?", input.ID).Find(&Lecture)

	if getLecture.RowsAffected < 1 {
		errorCode <- "RESULT_STUDENT_NOT_FOUND_404"
		return &Lecture, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Lecture, <-errorCode

}
