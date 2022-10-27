package updateUser

//type InputUpdateUser struct {
//	ID      int    `validate:"required,uuid"`
//	Name    string `json:"name" validate:"required"`
//	Phone   string `json:"phone" validate:"required"`
//	Email   string `json:"email" validate:"required,email"`
//	Address string `json:"address" validate:"required"`
//	Active  bool   `json:"active" validate:"required"`
//	Role    string `json:"role" validate:"required"`
//}

type InputUpdateUser struct {
	ID         int    `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Phone      string `json:"phone" binding:"required"`
	ProfileURL string `json:"profile_url"`
	Active     bool   `json:"active"`
	Role       string `json:"role" binding:"required"`
	Address    string `json:"address"`
}
