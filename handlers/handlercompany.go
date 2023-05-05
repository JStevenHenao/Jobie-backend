package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tidy/db"
	"tidy/dtoRequest"
	"tidy/models"

	"github.com/gorilla/mux"
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

func GetCompany(rw http.ResponseWriter, r *http.Request) {

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

	vars := mux.Vars(r)
	idCompany, _ := strconv.Atoi(vars["id"])
	companies := make([]models.Company, 0)
	//db.Database.Model(&models.Company{}).Select("users.name, emails.email").Joins("inner join company_services on companies.id = company_services.id_company").Scan(&result{})
	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

	rows, err := db.Database.Table("companies").Select("companies.id,companies.nit,companies.name,companies.calification,companies.image,companies.phone_number,companies.state,companies.reason").Joins("inner join company_services on companies.id = company_services.id_company").Where("company_services.id_service = ?", idCompany).Rows()
	var company models.Company

	if err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {

		for rows.Next() {

			rows.Scan(&company.Id, &company.Name, &company.NIT, &company.Calification, &company.Image, &company.PhoneNumber, &company.State, &company.Reason)
			companies = append(companies, company)

		}
		sendData(rw, companies, http.StatusCreated)
	}
}
