package models

import "time"

type Reference struct {
	Id          int
	UserName    string
	Description string
	Date        time.Time
	Stars       int
	IdCompany   int
	Company     Company `gorm:"foreignKey:IdCompany"`
}
