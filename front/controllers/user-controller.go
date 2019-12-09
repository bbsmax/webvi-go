package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	dto2 "webvi-go/front/dto"
	"webvi-go/front/services"
)

type UserController struct {
	DB *gorm.DB
}

var (
	userService = services.UserService{}
)

//회원로그인
func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login : ")
}

//회원로그아웃
func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login : ")
}

//회원정보찾기
func (c *UserController) Search(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login : ")
}

//회원가입
func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup")
	//userService := services.UserService{}
	if err := r.ParseForm(); err != nil {
		//TODO 에러메세지 발생.
	}

	requestData := &dto2.UserRequest{}
	//get으로 넘어온 변수는 schema
	if err := json.NewDecoder(r.Body).Decode(requestData); err != nil {
		//TODO 에러메세지 발생.
	}

	//validate
	userService.Signup(requestData, c.DB)
}

//회원정보수정
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {

}

//회원삭제
