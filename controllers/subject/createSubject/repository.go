package createSubject

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateSubjectRepository(input *model.Subject) (*model.Subject, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateSubjectRepository(input *model.Subject) (*model.Subject, string) {

	var Subjects model.Subject
	db := r.db.Model(&Subjects)
	errorCode := make(chan string, 1)

	checkSubjectExist := db.Select("*").Where("name = ?", input.Name).Find(&Subjects)

	if checkSubjectExist.RowsAffected > 0 {
		errorCode <- "CREATE_Subject_CONFLICT_409"
		return &Subjects, <-errorCode
	}

	Subjects.Name = input.Name
	Subjects.TeacherId = input.TeacherId
	Subjects.LectureId = input.LectureId

	addNewSubject := db.Create(&Subjects)
	db.Commit()

	if addNewSubject.Error != nil {
		errorCode <- "CREATE_Subject_FAILED_403"
		return &Subjects, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Subjects, <-errorCode
}
