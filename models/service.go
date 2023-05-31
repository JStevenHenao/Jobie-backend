package models

import (
	"tidy/db"
)

type Service struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	NumConsult  int64  `json:"numConsult"`
	Icon        string ` json:"icon"`
}

func MigrateModel() {

	db.Database.AutoMigrate(CompanyService{})

}
