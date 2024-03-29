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
	errorMsg = &utils.ReturnMessage{}
)

type UserDao struct{}

func (u *UserDao) Login(db *gorm.DB, requestData *dto.LoginRequest) (*dto.UserResponse, error) {
	user := User{}
	res := db.Where("email = ? AND password = ?", requestData.Email, requestData.Password).First(&user)

	if res.RecordNotFound() {
		return nil, fmt.Errorf("RecordNotFound")
	} else if err := res.Error; err != nil {
		return nil, err
	}

	userResponse := &dto.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		Role:   user.Role,
		Avator: user.Avator,
	}

	return userResponse, nil
}

//이메일로 회원유무를 검사.
func (u *UserDao) FindByEmail(db *gorm.DB, email string) (bool, *utils.ReturnMessage) {

	user := User{}

	db.LogMode(true)
	query := db.Where("email = ?", email).Find(&user)

	if res := query.RecordNotFound(); res {
		return false, errorMsg.ReturnMsg(fmt.Sprintf("user email [%v] not find", email), http.StatusNoContent, "no content")
	}

	if err := query.Error; err != nil {
		return false, errorMsg.ReturnMsg(err.Error(), http.StatusInternalServerError, "internal server error")
	}

	return true, nil
}

func (u *UserDao) FindByID(db *gorm.DB, ID string) (bool, *utils.ReturnMessage) {

	user := User{}

	db.LogMode(true)
	query := db.Where("ID = ?", ID).Find(&user)

	if res := query.RecordNotFound(); res {
		return false, errorMsg.ReturnMsg(fmt.Sprintf("user ID [%v] contents not find", ID), http.StatusNoContent, "no content")
	}

	if err := query.Error; err != nil {
		return false, errorMsg.ReturnMsg(err.Error(), http.StatusInternalServerError, "internal server error")
	}

	return true, nil
}

//회원가입 진행.
func (u *UserDao) Create(db *gorm.DB, requestData *dto.UserRequest) *utils.ReturnMessage {

	user := User{
		ID:       requestData.ID,
		Password: requestData.Password,
		Name:     requestData.Name,
		Email:    requestData.Email,
		Role:     "NORMAL",
	}

	res := db.Save(&user)

	if err := res.Error; err != nil {
		return errorMsg.ReturnMsg(err.Error(), http.StatusInternalServerError, "internal server error")
	}

	return nil
}

//회원정보 업데이트.
func (u *UserDao) Update(db *gorm.DB, ID string, updateData map[string]string) (*dto.UserResponse, *utils.ReturnMessage) {

	user := User{}
	res := db.Model(&user).Where("ID = ?", ID).Updates(updateData)

	if err := res.Error; err != nil {
		return nil, errorMsg.ReturnMsg(err.Error(), http.StatusInternalServerError, "internal server error")
	}

	res = db.Where("ID = ?", ID).First(&user)
	if err := res.Error; err != nil {
		return nil, errorMsg.ReturnMsg(err.Error(), http.StatusInternalServerError, "internal server error")
	}

	outData := &dto.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		Role:   user.Role,
		Avator: user.Avator,
	}
	return outData, nil
}

//회원 상세 정보.
func (u *UserDao) Get(db *gorm.DB, ID string) (*dto.UserResponse, *utils.ReturnMessage) {
	user := User{}
	db.LogMode(true)
	res := db.Where("ID = ?", ID).First(&user)
	if err := res.Error; err != nil {
		return nil, errorMsg.ReturnMsg(err.Error(), http.StatusInternalServerError, "internal server error")
	}

	outData := &dto.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		Role:   user.Role,
		Avator: user.Avator,
	}
	return outData, nil
}

//회원정보 삭제.
func (u *UserDao) Delete(db *gorm.DB, ID string) (string, *utils.ReturnMessage) {

	res := db.Where("id = ?", ID).Delete(&User{})

	if err := res.Error; err != nil {
		return "", errorMsg.ReturnMsg(err.Error(), http.StatusInternalServerError, "internal server error")
	}

	return fmt.Sprintf("deleted ID [%v]", ID), nil
}
