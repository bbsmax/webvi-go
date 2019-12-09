package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"webvi-go/front/dto"
	user2 "webvi-go/front/model/user"
)

type UserService struct{}

var (
	//UserDto 객체생성
	userDto = user2.UserDto{}
)

func (u *UserService) Signup(requestData *dto.UserRequest, db *gorm.DB) (bool, error) {
	fmt.Println("requestData : ", requestData)

	//1. 이메일로 회원이 존재하는지 검사.
	return userDto.FindUser(requestData, db)
	//2. 회원비밀번호를 sha256, sha512로 변환.
	//3. 회원데이터 저장

}
