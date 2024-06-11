package models

type Bank struct {
	Bank_code string ` gorm:"primaryKey" `
	Address   string
	Name      string
}

func (a *Bank) TableName() string {
	return "bank"
}
