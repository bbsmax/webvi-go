package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"webvi-go/controllers/users"

	"github.com/pelletier/go-toml"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	SERVER_PORT       = ":8080"
	DATABASE_URL      = ""
	DATABASE_USERNAME = ""
	DATABASE_PASSWORD = ""
	DATABASE_SCHEME   = ""
)

func StartApp() {
	config, err := toml.LoadFile("webvi-go/config.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	fmt.Println("config : ", config.Get("database.username").(string))

	serverStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&serverTimezone=Asia/Seoul&loc=local",
		config.Get("database.username").(string),
		config.Get("database.password").(string),
		config.Get("database.server").(string),
		config.Get("database.port").(string),
		config.Get("database.schema").(string))

	db, err := gorm.Open("mysql", serverStr)

	if err != nil {
		log.Panicf("error : [%v]", err)
		panic(err)
	}

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/user", users.Signup).Methods(http.MethodPost)

	if err := http.ListenAndServe(SERVER_PORT, router); err != nil {
		log.Panicf("Server Start Error:[%v]", err)
	}
}
