package createLecture

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateLectureRepository(input *model.Lecture) (*model.Lecture, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateLectureRepository(input *model.Lecture) (*model.Lecture, string) {

	var Lectures model.Lecture
	db := r.db.Model(&Lectures)
	errorCode := make(chan string, 1)

	checkLectureExist := db.Select("*").Where("name = ?", input.Name).Find(&Lectures)

	if checkLectureExist.RowsAffected > 0 {
		errorCode <- "CREATE_Lecture_CONFLICT_409"
		return &Lectures, <-errorCode
	}

	Lectures.Name = input.Name
	Lectures.TutorId = input.TutorId

	addNewLecture := db.Debug().Create(&Lectures)
	db.Commit()

	if addNewLecture.Error != nil {
		errorCode <- "CREATE_LECTURE_FAILED_403"
		return &Lectures, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Lectures, <-errorCode
}
