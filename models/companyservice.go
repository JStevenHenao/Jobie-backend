package models

type CompanyService struct {
	Id        int64
	IdCompany int64
	Company   Company `gorm:"foreignKey:IdCompany"`
	IdClient  int64
	Client    Client `gorm:"foreignKey:IdClient"`
}
