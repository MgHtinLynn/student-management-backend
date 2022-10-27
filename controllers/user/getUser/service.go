package getUser

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type Service interface {
	GetUserService(input *InputGetUser) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceUser(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetUserService(input *InputGetUser) (*model.User, string) {

	var user = model.User{
		ID: input.ID,
	}
	getUser, errCreateStudent := s.repository.UserRepository(&user)

	return getUser, errCreateStudent
}
