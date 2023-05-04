package models

type Client struct {
	Id          int64
	Name        string
	LastName    string
	Profession  string
	TypeId      string
	Ide         string
	NumberPhone int
	Email       string
	Password    string
	UserName    string
	IdCompany   int
	Company     Company `gorm:"foreignKey:IdCompany"`
}
