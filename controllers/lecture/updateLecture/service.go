package updateLecture

import model "github.com/MgHtinLynn/final-year-project-mcc/service/models"

type Service interface {
	UpdateLectureService(input *InputUpdateLecture) (*model.Lecture, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateLectureService(input *InputUpdateLecture) (*model.Lecture, string) {

	lectures := model.Lecture{
		ID:      input.ID,
		Name:    input.Name,
		TutorId: input.TutorId,
	}

	resultUpdateLecture, errCreateLecture := s.repository.UpdateLectureRepository(&lectures)

	return resultUpdateLecture, errCreateLecture
}
