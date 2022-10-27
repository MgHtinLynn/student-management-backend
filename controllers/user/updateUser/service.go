package updateUser

import model "github.com/MgHtinLynn/final-year-project-mcc/service/models"

type Service interface {
	UpdateUserService(input *InputUpdateUser) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateUserService(input *InputUpdateUser) (*model.User, string) {

	Users := model.User{
		ID:         input.ID,
		Name:       input.Name,
		Email:      input.Email,
		Phone:      input.Phone,
		Active:     input.Active,
		Address:    input.Address,
		Role:       input.Role,
		ProfileUrl: input.ProfileURL,
	}

	resultUpdateUser, errCreateUser := s.repository.UpdateUserRepository(&Users)

	return resultUpdateUser, errCreateUser
}
