package createUser

type InputCreateUser struct {
	Name       string `json:"name" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Address    string `json:"address" binding:"required"`
	Active     bool   `json:"active"`
	ProfileURL string `json:"profile_url"`
	Role       string `json:"role" binding:"required"`
}
