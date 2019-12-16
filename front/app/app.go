package app

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/pelletier/go-toml"
	"os"
	"webvi-go/front/dto"
	"webvi-go/front/model/user"

	"log"
	"net/http"
	"webvi-go/front/controllers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	cache *redis.Client
)

func StartApp() {

	tomlFile, err := os.Open("webvi-go/front/config.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	defer tomlFile.Close()
	config := dto.Config{}
	if err := toml.NewDecoder(tomlFile).Decode(&config); err != nil {
		panic(err)
	}

	serverConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Schema)

	db, err := gorm.Open("mysql", serverConfig)

	if err != nil {
		log.Panicf("error : [%v]", err)
	}

	db.AutoMigrate(&user.User{})

	defer db.Close()

	initCache(config)

	aaa := struct {
		ID       string
		Password string
	}{ID: "aaaaaa", Password: "bbbb"}

	fmt.Println("aaa : ", aaa)
	test, _ := json.Marshal(aaa)
	fmt.Println("test", string(test))
	pong := cache.Do("SETEX", "bb", "120", string(test))
	//cache.Set("name", "Elliot", 0)
	fmt.Println("pong : ", pong)

	userController := controllers.UserController{
		DB: db,
	}
	router := mux.NewRouter()
	router.HandleFunc("/login", userController.Login).Methods(http.MethodPost)
	router.HandleFunc("/login", userController.Logout).Methods(http.MethodGet)
	router.HandleFunc("/user", userController.Create).Methods(http.MethodPost)
	router.HandleFunc("/user/{ID}", userController.Update).Methods(http.MethodPatch)
	router.HandleFunc("/user/{ID}", userController.Get).Methods(http.MethodGet)
	router.HandleFunc("/user/{ID}", userController.Delete).Methods(http.MethodDelete)

	fmt.Printf("Webvi-go Server port[%v] Start!! ", config.Webserver.Port)

	if err := http.ListenAndServe(config.Webserver.Port, router); err != nil {
		log.Printf("Server Start Error:[%v]", err)
	}
}

func initCache(config dto.Config) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Password,
	})

	cache = client
}
