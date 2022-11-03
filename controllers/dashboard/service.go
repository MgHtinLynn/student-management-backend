package getDashboard

import (
	model "github.com/MgHtinLynn/final-year-project-mcc/service/models"
)

type dataDashboard struct {
	//user        *model.User
	count       *int64
	activeCount *int64
}

type Service interface {
	GetDashboardService(input *InputGetUser) (*model.User, *int64, *int64, string)
}

type service struct {
	repository Repository
}

func NewServiceDashboard(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetDashboardService(input *InputGetUser) (*model.User, *int64, *int64, string) {

	var user = model.User{
		ID: input.ID,
	}

	getUser, count, activeCount, errUser := s.repository.DashboardRepository(&user)

	return getUser, count, activeCount, errUser
}
