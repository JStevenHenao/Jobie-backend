package models

type SocialNetwork struct {
	Id       int64
	Type     string
	Link     string
	IdClient int64
	Client   Client `gorm:"foreignKey:IdClient"`
}
