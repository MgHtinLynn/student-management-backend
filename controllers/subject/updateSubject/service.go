package updateSubject

import model "github.com/MgHtinLynn/final-year-project-mcc/service/models"

type Service interface {
	UpdateSubjectService(input *InputUpdateSubject) (*model.Subject, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateSubjectService(input *InputUpdateSubject) (*model.Subject, string) {

	Subjects := model.Subject{
		ID:        input.ID,
		Name:      input.Name,
		TeacherId: input.TeacherId,
		LectureId: input.LectureId,
	}

	resultUpdateSubject, errCreateSubject := s.repository.UpdateSubjectRepository(&Subjects)

	return resultUpdateSubject, errCreateSubject
}
