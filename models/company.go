package models

type Company struct {
	Id           int64  `json:"id"`
	NIT          string `json:"nit"`
	Name         string `json:"name"`
	Calification int    `json:"calification"`
	Image        string `json:"image"`
	PhoneNumber  int    `json:"phoneNumber"`
	State        string `json:"state"`
}
