package models

type CompanyService struct {
	Id        int64
	IdCompany int64
	Company   Company `gorm:"foreignKey:IdCompany"`
	IdService int64
	Service   Service `gorm:"foreignKey:IdService"`
}
