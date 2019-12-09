package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"webvi-go/front/dto"
)

//DB와 데이터를 주고 받는 기능을 한다.

type UserDto struct{}

//이메일로 회원유무를 검사.
func (u *UserDto) FindUser(requestData *dto.UserRequest, db *gorm.DB) (bool, error) {
	userEmail := requestData.Email

	db.LogMode(true)
	if err := db.Where("email = ?", userEmail); err != nil {

	}

	return true, nil
}

//회원가입 진행.
func (u *UserDto) Signup(requestData *dto.UserRequest) {
	fmt.Println("userDto : ", requestData)
}
