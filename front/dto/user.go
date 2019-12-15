package dto

import (
	"errors"
	"fmt"
	"webvi-go/front/utils"
)

type LoginRequest struct {
	Email    string `json:"email""`
	Password string `json:"password"`
}

type LoginResponse struct {
	LoginSession string `json:"loginSession"`
}

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

func (r LoginRequest) Validate() error {
	var errs error

	if !utils.EmailCheck(r.Email) {
		errs = errors.New("The email field should be a valid email address!")
	} else if !utils.PasswordCheck(r.Password) {
		errs = errors.New("Password must be 8 characters or more, special characters, numbers, and letters.")
	}
	return errs
}

func (r UserRequest) Validate() error {
	var errs error
	if !utils.EmailCheck(r.Email) {
		errs = errors.New("The email field should be a valid email address!")
	} else if !utils.PasswordCheck(r.Password) {
		errs = errors.New("Password must be 8 characters or more, special characters, numbers, and letters.")
	}

	if ok, count := utils.NameCheck(r.Name); !ok {
		errs = fmt.Errorf("The name must be at least %v characters long.", count)
	}

	return errs
}
