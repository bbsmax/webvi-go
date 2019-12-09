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

	if res := query.RecordNotFound(); res {
		return false, errorMsg.ReturnErrorMsg(fmt.Sprintf("user email [%v] not find", userEmail), http.StatusNoContent)
	}

	if err := query.Error; err != nil {
		return false, errorMsg.ReturnErrorMsg(err.Error(), http.StatusInternalServerError)
	}

	return true, nil
}

//회원가입 진행.
func (u *UserDto) SignupCreate(requestData *dto.UserRequest, db *gorm.DB) *utils.ErrorMessage {

	user := User{
		ID:       requestData.ID,
		Password: requestData.Password,
		Name:     requestData.Name,
		Email:    requestData.Email,
		Role:     "BUSINESS_NORMAL",
	}

	res := db.Save(&user)

	if err := res.Error; err != nil {
		return errorMsg.ReturnErrorMsg(err.Error(), http.StatusInternalServerError)
	}

	return nil
}

//회원정보 업데이트.
func (u *UserDto) SignupUpdate(requestData *dto.UserRequest) {

}

//회원가입 진행.
func (u *UserDto) SignupDelete(requestData *dto.UserRequest) {
	fmt.Println("userDto : ", requestData)
}
