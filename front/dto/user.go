package dto

import (
	"errors"
	"webvi-go/front/utils"
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
	} else if r.Password == "" {
		errs = errors.New("The password field should be a required!")
	} else if r.Name == "" {
		errs = errors.New("The name field should be a required!")
	}
	return errs
}
