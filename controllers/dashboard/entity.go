package getDashboard

import model "github.com/MgHtinLynn/final-year-project-mcc/service/models"

type InputGetUser struct {
	ID int `validate:"required"`
}

type DataDashboard struct {
	user        *model.User
	count       *int64
	activeCount *int64
}
