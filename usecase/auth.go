package usecase

type Login struct{}

type LoginInterface interface {
	Authenticate(Username, Password string) bool
}

func NewLogin() LoginInterface {
	return &Login{}
}

func (masuk *Login) Authenticate(username, password string) bool {
	if username == "Admin" && password == "admin123" {
		return true
	}

	return false
}
