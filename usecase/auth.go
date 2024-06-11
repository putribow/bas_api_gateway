package usecase

import (
	"TaskGo/models"
	"TaskGo/utils"
)

type Login struct{}

type LoginInterface interface {
	Authenticate(Username, Password string) bool
}

func NewLogin() LoginInterface {
	return &Login{}
}

func (masuk *Login) Authenticate(username, password string) bool {

	accounts := models.Account{}
	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()
	defer db.Close()

	orm.Find(&accounts, "username = ? AND password = ? ", username, password)
	if accounts.AccountID == "" {

		return false
	}

	return true
}
