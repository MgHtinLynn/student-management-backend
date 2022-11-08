package updatePassword

import model "github.com/MgHtinLynn/final-year-project-mcc/service/models"

type Service interface {
	UpdateUserService(input *InputUpdatePassword) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateUserService(input *InputUpdatePassword) (*model.User, string) {

	Users := model.User{
		ID:       input.ID,
		Password: input.Password,
	}

	resultUpdateUser, errCreateUser := s.repository.UpdateUserRepository(&Users)

	return resultUpdateUser, errCreateUser
}
