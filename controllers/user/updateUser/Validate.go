package updateUser

// User contains user information
type User struct {
	ID        int    `validate:"required"`
	Name      string `validate:"required"`
	Phone     string `validate:"required"`
	Email     string `validate:"required,email,unique"`
	Addresses string `validate:"required,dive,required"`
}
