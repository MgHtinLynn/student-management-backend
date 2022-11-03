package createSubject

import model "github.com/MgHtinLynn/final-year-project-mcc/service/models"

type Service interface {
	CreateSubjectService(input *InputCreateSubject) (*model.Subject, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateSubjectService(input *InputCreateSubject) (*model.Subject, string) {

	Subjects := model.Subject{
		Name:      input.Name,
		TeacherId: input.TeacherId,
		LectureId: input.LectureId,
	}

	resultCreateUser, errCreateUser := s.repository.CreateSubjectRepository(&Subjects)

	return resultCreateUser, errCreateUser
}
