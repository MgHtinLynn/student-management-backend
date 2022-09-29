package loginAuth

import (
	"fmt"
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type Service interface {
	LoginService(input *InputLogin) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *InputLogin) (*model.User, string) {

	user := model.User{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	fmt.Println(resultLogin)
	fmt.Println(errLogin)
	return resultLogin, errLogin
}
