package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tidy/db"
	"tidy/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetService(rw http.ResponseWriter, r *http.Request) {

	services := []models.Service{}
	db.Database.Find(&services)
	sendData(rw, services, http.StatusOK)
}

func CreateService(rw http.ResponseWriter, r *http.Request) {
	service := models.Service{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&service); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&service)
		sendData(rw, service, http.StatusCreated)
	}
}

func UpdateService(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro
	var userId int64

	if service_ant, err := getServiceById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		userId = service_ant.Id

		user := models.Service{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&user); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			db.Database.Save(&user)
			sendData(rw, user, http.StatusOK)
		}
	}

}

func getServiceById(r *http.Request) (models.Service, *gorm.DB) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.Service{}

	if err := db.Database.First(&user, userId); err.Error != nil {
		return user, err
	} else {
		return user, nil
	}
}
