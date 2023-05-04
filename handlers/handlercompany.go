package handlers

import (
	"encoding/json"
	"net/http"
	"tidy/db"
	"tidy/models"
)

func CreateCompany(rw http.ResponseWriter, r *http.Request) {
	company := models.Company{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&company); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&company)
		sendData(rw, company, http.StatusCreated)
	}
}
