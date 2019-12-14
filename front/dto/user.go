package dto

import (
	"errors"
	"webvi-go/front/utils"
)

const (
	passwordPattern = "/^(?=.*[a-zA-Z])(?=.*[!@#$%^*+=-])(?=.*[0-9]).{6,16}$/" //영문 + 특수문자 + 숫자정규식
)

type UserRequest struct {
	ID       string `json:"ID"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Avator   string `json:"avator"`
}

type UserResponse struct {
	ID     string `json:"ID"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`
	Avator string `json:"avator"`
}

func (r UserRequest) Validate() error {
	var errs error
	if !utils.EmailCheck(r.Email) {
		errs = errors.New("The email field should be a valid email address!")
	} else if !utils.PasswordCheck(r.Password) {
		errs = errors.New("Password must be 8 characters or more, special characters, numbers, and letters.")
	} else if !utils.NameCheck(r.Name) {
		errs = errors.New("PThe name must be at least 4 characters long.")
	}
	return errs
}
