package models

type Account struct {
	AccountID string ` gorm:"primaryKey" `
	Username  string ` gorm:"username" `
	Password  string
	Name      string
}

func (a *Account) TableName() string {
	return "account"
}
