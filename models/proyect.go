package models

type Proyect struct {
	Id          int64
	Name        string
	Description string
	IdCompany   int64
	Company     Company `gorm:"foreignKey:IdCompany"`
}
