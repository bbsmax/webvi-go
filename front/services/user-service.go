package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"net/http"
	"webvi-go/front/dto"
	"webvi-go/front/model/user"
	"webvi-go/front/utils"
)

type UserService struct{}

var (
	//UserDto 객체생성
	userDto  = user.UserDto{}
	errorMsg = utils.ErrorMessage{}
)

func (u *UserService) Create(requestData *dto.UserRequest, db *gorm.DB) (bool, *utils.ErrorMessage) {

	//1. 이메일로 회원이 존재하는지 검사.
	if isUser, err := userDto.FindUser(requestData, db); !isUser {
		//회원이 존재하지 않은 경우.
		//2. 회원비밀번호를 sha256, sha512로 변환.
		id := uuid.New()
		password := requestData.Password
		requestData.ID = id.String()
		requestData.Password = password
		//3. 회원데이터 저장
		userDto.SignupCreate(requestData)
	} else if isUser {
		//회원이 존재하는 경우.
		return isUser, errorMsg.ReturnErrorMsg(fmt.Sprintf("user email [%v] is exist", requestData.Email), http.StatusInternalServerError)
	} else if err != nil {
		//에러가 발생한 경우.
		return isUser, err
	}

	return true, nil
}
