package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"webvi-go/front/dto"
	"webvi-go/front/utils"
)

//DB와 데이터를 주고 받는 기능을 한다.

var (
	errorMsg = &utils.ErrorMessage{}
)

type UserDto struct{}

//이메일로 회원유무를 검사.
func (u *UserDto) FindUser(requestData *dto.UserRequest, db *gorm.DB) (bool, *utils.ErrorMessage) {
	userEmail := requestData.Email
	user := User{}

	db.LogMode(true)
	query := db.Where("email = ?", userEmail).Find(&user)

	if err := query.Error; err != nil {
		return false, errorMsg.ErrorMsg("user find id", http.StatusConflict, fmt.Errorf("[%v]", err))
	}

	if res := query.RecordNotFound(); res {
		return false, errorMsg.ErrorMsg("user find id", http.StatusNoContent, fmt.Errorf("[%v]", "not content"))
	}

	return true, nil
}

//회원가입 진행.
func (u *UserDto) Signup(requestData *dto.UserRequest) {
	fmt.Println("userDto : ", requestData)
}
