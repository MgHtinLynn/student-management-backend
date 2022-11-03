package createLecture

import model "github.com/MgHtinLynn/final-year-project-mcc/service/models"

type Service interface {
	CreateLectureService(input *InputCreateLecture) (*model.Lecture, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateLectureService(input *InputCreateLecture) (*model.Lecture, string) {

	lectures := model.Lecture{
		Name:    input.Name,
		TutorId: input.TutorId,
	}

	resultCreateUser, errCreateUser := s.repository.CreateLectureRepository(&lectures)

	return resultCreateUser, errCreateUser
}
