package main

import (
	"log"
	"net/http"
	"tidy/handlers"

	"github.com/gorilla/mux"
)

func main() {

	//models.MigrateModel()
	mux := mux.NewRouter()

	mux.HandleFunc("/api/jobie/services", handlers.GetService).Methods("GET")
	mux.HandleFunc("/api/jobie/services", handlers.CreateService).Methods("POST")
	mux.HandleFunc("/api/jobie/services/{id:[0-9]+}", handlers.UpdateService).Methods("PUT")

	mux.HandleFunc("/api/jobie/company", handlers.CreateCompany).Methods("POST")
	mux.HandleFunc("/api/jobie/services/company", handlers.CreateServicesCompany).Methods("POST")

	mux.HandleFunc("/api/jobie/services/company/{id:[0-9]+}", handlers.GetServicesForCompany).Methods("GET")

	/*mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")

	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")*/

	log.Fatal(http.ListenAndServe(":3000", mux))
}
