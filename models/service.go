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
	db.Database.AutoMigrate(Service{})
	db.Database.AutoMigrate(Company{})
	db.Database.AutoMigrate(Client{})
	db.Database.AutoMigrate(Proyect{})
	db.Database.AutoMigrate(CompanyService{})
	db.Database.AutoMigrate(Reference{})
	db.Database.AutoMigrate(SocialNetwork{})

}
