package updatePassword

type InputUpdatePassword struct {
	ID       int    `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
