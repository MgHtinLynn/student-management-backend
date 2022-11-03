package updateLecture

import (
	"fmt"
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateLectureRepository(input *model.Lecture) (*model.Lecture, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateLectureRepository(input *model.Lecture) (*model.Lecture, string) {

	var Lecture model.Lecture
	db := r.db.Model(&Lecture)
	errorCode := make(chan string, 1)

	checkEmailExist := db.Select("*").Not("id = ?", input.ID).Where("name = ?", input.Name).Find(&Lecture)

	if checkEmailExist.RowsAffected > 0 {
		errorCode <- "UPDATE_Lecture_EMAIL_CONFLICT_400"
		return &Lecture, <-errorCode
	}

	Lecture.ID = input.ID

	checkLectureId := db.Select("*").Where("id = ?", input.ID).Find(&Lecture)

	if checkLectureId.RowsAffected < 1 {
		errorCode <- "UPDATE_Lecture_NOT_FOUND_404"
		return &Lecture, <-errorCode
	}

	Lecture.Name = input.Name
	Lecture.TutorId = input.TutorId

	fmt.Println("Lecture", Lecture)

	updateLecture := db.Select("*").Where("id = ?", input.ID).Updates(Lecture)

	fmt.Println("updateLecture", updateLecture)

	if updateLecture.Error != nil {
		errorCode <- "UPDATE_Lecture_FAILED_403"
		return &Lecture, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Lecture, <-errorCode
}
