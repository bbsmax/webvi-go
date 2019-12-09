package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"webvi-go/front/dto"
	user2 "webvi-go/front/model/user"
	"webvi-go/front/utils"
)

type UserService struct{}

var (
	//UserDto 객체생성
	userDto = user2.UserDto{}
)

func (u *UserService) Signup(requestData *dto.UserRequest, db *gorm.DB) (bool, *utils.ErrorMessage) {
	fmt.Println("requestData : ", requestData)

	//1. 이메일로 회원이 존재하는지 검사.
	if isUser, err := userDto.FindUser(requestData, db); !isUser {
		//회원이 존재하지 않은 경우.
		//2. 회원비밀번호를 sha256, sha512로 변환.
		//3. 회원데이터 저장
	} else if isUser {
		//회원이 존재하는 경우.
	} else if err != nil {
		//에러가 발생한 경우.
	}

	return true, nil
}
