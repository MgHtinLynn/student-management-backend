package auth

type LoginService interface {
	LoginService(email string, password string) bool
}

type loginInfo struct {
	email    string
	password string
}

func loginService() *loginInfo {
	return &loginInfo{
		email:    "htinlin01@gmail.com",
		password: "password",
	}
}

func (info *loginInfo) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}
