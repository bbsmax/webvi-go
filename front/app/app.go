package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/pelletier/go-toml"
	user2 "webvi-go/front/model/user"

	"log"
	"net/http"
	"webvi-go/front/controllers"

	_ "github.com/pelletier/go-toml"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func StartApp() {
	config, err := toml.LoadFile("webvi-go/front/config.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	serverConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Get("database.username").(string),
		config.Get("database.password").(string),
		config.Get("database.host").(string),
		config.Get("database.port").(string),
		config.Get("database.schema").(string))

	db, err := gorm.Open("mysql", serverConfig)

	if err != nil {
		log.Panicf("error : [%v]", err)
	}

	db.AutoMigrate(&user2.User{})

	defer db.Close()
	//userController := controller.UserController{}
	userController := controllers.UserController{
		DB: db,
	}
	router := mux.NewRouter()
	router.HandleFunc("/login", userController.Login).Methods(http.MethodGet)
	router.HandleFunc("/login", userController.Logout).Methods(http.MethodGet)
	router.HandleFunc("/user", userController.Create).Methods(http.MethodPost)
	router.HandleFunc("/user/{ID}", userController.Update).Methods(http.MethodPatch)
	router.HandleFunc("/user/{ID}", userController.Get).Methods(http.MethodGet)
	router.HandleFunc("/user/{ID}", userController.Delete).Methods(http.MethodDelete)

	if err := http.ListenAndServe(config.Get("webserver.port").(string), router); err != nil {
		log.Printf("Server Start Error:[%v]", err)
	}

}
