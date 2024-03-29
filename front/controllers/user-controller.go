package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jinzhu/gorm"

	"net/http"
	"time"
	"webvi-go/front/dto"
	"webvi-go/front/services"
	"webvi-go/front/utils"

	"github.com/go-redis/redis"
)

type UserController struct {
	DB     *gorm.DB
	Client *redis.Client
}

var (
	userService = services.UserService{}
	response    = utils.ReturnMessage{}
)

//회원로그인
func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login : ")
	if r.Method != http.MethodPost {
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	if err := r.ParseForm(); err != nil {
		//TODO 에러메세지 발생.
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	requestData := &dto.LoginRequest{}

	if err := json.NewDecoder(r.Body).Decode(requestData); err != nil {
		//TODO 에러메세지 발생.
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	//validate 체크.
	if err := requestData.Validate(); err != nil {
		response.ResponseMsg(w, err.Error(), http.StatusBadRequest, "bad request")
		return
	}

	token, err := userService.Login(c.DB, c.Client, requestData)

	if err != nil {
		response.ResponseMsg(w, err.Message, http.StatusInternalServerError, "user not exist")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "x-access-token",
		Value:   token.JwtTokenString,
		Expires: token.CookieExpiredTime,
	})

	fmt.Println("Hello!!!!!!!!!!!")

}

func (c *UserController) Welcome(w http.ResponseWriter, r *http.Request) {
	redisClient := c.Client
	cookie, err := r.Cookie("x-access-token")

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := cookie.Value

	token, error := jwt.Parse(sessionToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("AllYourBase"), nil
	})

	//fmt.Println("Hello", token, error)

	if error != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error : %+v", error.Error()))
		return
	}

	if token.Valid {
		fmt.Println("token 성공")
		json.NewEncoder(w).Encode("Success authorization token")
		return
	} else {
		json.NewEncoder(w).Encode("Invalid authorization token")
		return
	}

	response := redisClient.Do("GET", sessionToken)
	fmt.Println("response.val", response.Val())
	if response == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	responseData := dto.UserResponse{}
	fmt.Println("type :", reflect.TypeOf(response.Val()))
	json.Unmarshal([]byte(response.Val().(string)), &responseData)

	w.Write([]byte(fmt.Sprintf("Welcome %s!", responseData.Name)))
}

func (c *UserController) Refresh(w http.ResponseWriter, r *http.Request) {
	redisClient := c.Client
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {

			//w.WriteHeader(http.Status)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := cookie.Value

	response := redisClient.Do("GET", sessionToken)

	if response == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// (END) The code uptil this point is the same as the first part of the `Welcome` route

	// Now, create a new session token for the current user
	newSessionToken := uuid.New().String()
	_ = redisClient.Do("SETEX", newSessionToken, "120", response.Val())

	// Delete the older session token
	_ = redisClient.Do("DEL", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(30 * time.Second),
	})
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

	if r.Method != http.MethodPost {
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	if err := r.ParseForm(); err != nil {
		//TODO 에러메세지 발생.
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	requestData := &dto.UserRequest{}
	//get으로 넘어온 변수는 schema r.Form
	//post r.PostForm
	//body r.Body
	if err := json.NewDecoder(r.Body).Decode(requestData); err != nil {
		//TODO 에러메세지 발생.
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	//validate
	if err := requestData.Validate(); err != nil {
		response.ResponseMsg(w, err.Error(), http.StatusBadRequest, "bad request")
		return
	}

	if _, err := userService.Create(c.DB, requestData); err != nil {
		response.ResponseMsg(w, err.Message, err.Code, err.Status)
	}
}

//상세 회원정보
func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	if err := r.ParseForm(); err != nil {
		//TODO 에러메세지 발생.
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	param := mux.Vars(r)
	var ID string
	if id, err := uuid.Parse(param["ID"]); err != nil {
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	} else {
		ID = id.String()
	}

	userData, err := userService.Get(c.DB, ID)

	if err != nil {
		response.ResponseMsg(w, err.Message, err.Code, err.Status)
		return
	}

	response.ResponseData(w, userData, http.StatusOK, "OK")
}

//회원정보 업데이트
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPatch {
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		//TODO 에러메세지 발생.
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	param := mux.Vars(r)
	var ID string
	if id, err := uuid.Parse(param["ID"]); err != nil {
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
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
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	requestData.ID = ID

	if updateData, err := userService.Update(c.DB, requestData, r); err != nil {
		response.ResponseMsg(w, err.Message, err.Code, err.Status)
	} else {
		response.ResponseData(w, updateData, http.StatusOK, "OK")
	}

}

//회원삭제
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	if err := r.ParseForm(); err != nil {
		//TODO 에러메세지 발생.
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	}

	param := mux.Vars(r)
	var ID string
	if id, err := uuid.Parse(param["ID"]); err != nil {
		response.ResponseMsg(w, "internal server error", http.StatusInternalServerError, "internal server error")
		return
	} else {
		ID = id.String()
	}

	deleteData, err := userService.Delete(c.DB, ID)

	if err != nil {
		response.ResponseMsg(w, err.Message, err.Code, err.Status)
		return
	}

	response.ResponseData(w, deleteData, http.StatusOK, "OK")
}
