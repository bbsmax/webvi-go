package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"webvi-go/controllers/users"
)

const (
	SERVER_PORT       = ":8080"
	DATABASE_URL      = ""
	DATABASE_USERNAME = ""
	DATABASE_PASSWORD = ""
	DATABASE_SCHEME   = ""
)

func StartApp() {
	router := mux.NewRouter()
	router.HandleFunc("/user", users.Signup).Methods(http.MethodPost)

	if err := http.ListenAndServe(SERVER_PORT, router); err != nil {
		log.Panicf("Server Start Error:[%v]", err)
	}
}
