package handlers

import (
	"encoding/json"
	"net/http"
	"tidy/db"
	"tidy/dtoRequest"
	"tidy/models"
)

func CreateCompany(rw http.ResponseWriter, r *http.Request) {

	///Agregar servicios
	company := models.Company{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&company); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&company)
		sendData(rw, company, http.StatusCreated)
	}
}

func CreateServicesCompany(rw http.ResponseWriter, r *http.Request) {
	companyService := dtoRequest.CompanyServiceRequest{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&companyService); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {

		for _, id := range companyService.GetIdServices() {

			companyServices := models.CompanyService{
				IdCompany: companyService.GetIdCompany(),
				IdService: id,
			}
			db.Database.Save(&companyServices)
		}
		sendData(rw, companyService, http.StatusCreated)
	}
}

func GetServicesForCompany(rw http.ResponseWriter, r *http.Request) {

	companyService := dtoRequest.CompanyServiceRequest{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&companyService); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {

		for _, id := range companyService.GetIdServices() {

			companyServices := models.CompanyService{
				IdCompany: companyService.GetIdCompany(),
				IdService: id,
			}
			db.Database.Save(&companyServices)
		}
		sendData(rw, companyService, http.StatusCreated)
	}
}
