package createUser

import model "github.com/MgHtinLynn/final-year-project-mcc/service/models"

type Service interface {
	CreateUserService(input *InputCreateUser) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateUserService(input *InputCreateUser) (*model.User, string) {

	Users := model.User{
		Name:       input.Name,
		Email:      input.Email,
		Phone:      input.Phone,
		Active:     input.Active,
		Address:    input.Address,
		Role:       input.Role,
		ProfileUrl: input.ProfileURL,
	}

	resultCreateUser, errCreateUser := s.repository.CreateUserRepository(&Users)

	return resultCreateUser, errCreateUser
}
