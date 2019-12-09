package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jinzhu/gorm"
	"net/http"
	"webvi-go/front/dto"
	"webvi-go/front/services"
	"webvi-go/front/utils"
)

type UserController struct {
	DB *gorm.DB
}

var (
	userService     = services.UserService{}
	responseMessage = utils.ReturnMessage{}
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

	if err := r.ParseForm(); err != nil {
		//TODO 에러메세지 발생.
		responseMessage.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	requestData := &dto.UserRequest{}
	//get으로 넘어온 변수는 schema r.Form
	//post r.PostForm
	//body r.Body
	if err := json.NewDecoder(r.Body).Decode(requestData); err != nil {
		//TODO 에러메세지 발생.
		responseMessage.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	//validate
	if err := requestData.Validate(); err != nil {
		responseMessage.ResponseMsg(w, err.Error(), http.StatusBadRequest, "bad request")
		return
	}

	if _, err := userService.Create(requestData, c.DB); err != nil {
		responseMessage.ResponseMsg(w, err.Message, err.Code, err.Status)
	}
}

//상세 회원정보
func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		//TODO 에러메세지 발생.
		responseMessage.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	param := mux.Vars(r)
	var ID string
	if id, err := uuid.Parse(param["ID"]); err != nil {
		responseMessage.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	} else {
		ID = id.String()
	}

	requestData := &dto.UserRequest{}
	//get으로 넘어온 변수는 schema r.Form
	//post r.PostForm
	//body r.Body
	decode := schema.NewDecoder()
	if err := decode.Decode(requestData, r.PostForm); err != nil {
		responseMessage.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	//validate
	if err := requestData.Validate(); err != nil {
		responseMessage.ResponseMsg(w, err.Error(), http.StatusBadRequest, "bad request")
		return
	}

	requestData.ID = ID

	if _, err := userService.Update(requestData, c.DB); err != nil {
		responseMessage.ResponseMsg(w, err.Message, err.Code, err.Status)
	}
}

//회원정보 업데이트
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		//TODO 에러메세지 발생.
		responseMessage.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	param := mux.Vars(r)
	var ID string
	if id, err := uuid.Parse(param["ID"]); err != nil {
		responseMessage.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	} else {
		ID = id.String()
	}

	requestData := &dto.UserRequest{}
	//get으로 넘어온 변수는 schema r.Form
	//post r.PostForm
	//body r.Body
	decode := schema.NewDecoder()
	if err := decode.Decode(requestData, r.PostForm); err != nil {
		responseMessage.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	//validate
	if err := requestData.Validate(); err != nil {
		responseMessage.ResponseMsg(w, err.Error(), http.StatusBadRequest, "bad request")
		return
	}

	requestData.ID = ID

	if _, err := userService.Update(requestData, c.DB); err != nil {
		responseMessage.ResponseMsg(w, err.Message, err.Code, err.Status)
	}
}

//회원삭제
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		//TODO 에러메세지 발생.
		responseMessage.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	param := mux.Vars(r)
	var ID string
	if id, err := uuid.Parse(param["ID"]); err != nil {
		responseMessage.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	} else {
		ID = id.String()
	}

	requestData := &dto.UserRequest{}
	//get으로 넘어온 변수는 schema r.Form
	//post r.PostForm
	//body r.Body
	decode := schema.NewDecoder()
	if err := decode.Decode(requestData, r.PostForm); err != nil {
		responseMessage.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	//validate
	if err := requestData.Validate(); err != nil {
		responseMessage.ResponseMsg(w, err.Error(), http.StatusBadRequest, "bad request")
		return
	}

	requestData.ID = ID

	if _, err := userService.Update(requestData, c.DB); err != nil {
		responseMessage.ResponseMsg(w, err.Message, err.Code, err.Status)
	}
}
