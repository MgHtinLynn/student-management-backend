package getLecture

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type Service interface {
	GetLectureService(input *InputGetLecture) (*model.Lecture, string)
}

type service struct {
	repository Repository
}

func NewServiceLecture(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetLectureService(input *InputGetLecture) (*model.Lecture, string) {

	var Lecture = model.Lecture{
		ID: input.ID,
	}
	getLecture, errCreateStudent := s.repository.LectureRepository(&Lecture)

	return getLecture, errCreateStudent
}
