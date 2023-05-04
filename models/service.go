package models

import (
	"tidy/db"
)

type Service struct {
	Id          int64
	Name        string
	Description string
	NumConsult  int64
	Icon        string
}

func MigrateModel() {

	db.Database.AutoMigrate(CompanyService{})

}
