package utils

import (
	"regexp"
	"unicode/utf8"
)

const (
	passwordLength = 8 //비밀번호 최소 길이
	nameUTFLength  = 2 //한글 이름 최소길이
	nameLength     = 4 //영문이릉 최소길이
)

func EmailCheck(email string) bool {
	var rxEmail = regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)
	isEmail := rxEmail.MatchString(email)
	return isEmail
}

//비밀번호은 영문조합 8자리 이상
func PasswordCheck(password string) bool {
	var rxPpassword = regexp.MustCompile(`/^(?=.*[a-zA-Z])(?=.*[!@#$%^*+=-])(?=.*[0-9]).{6,16}$/`) //영문 + 특수문자 + 숫자정규식

	if len(password) < passwordLength {
		return false
	}

	isPassword := rxPpassword.MatchString(password)
	return isPassword
}

//이름체크
func NameCheck(name string) bool {
	utfCount := utf8.RuneCountInString(name) //문자열의 실제 길이를 구함.
	normalCount := len(name)
	isNameCheck := true

	if normalCount == utfCount {
		//영문이름일 경우 4글자 이상
		if normalCount < nameLength {
			isNameCheck = false
		}
	} else {
		//한글이름일 경우 2글자 이상
		if utfCount < nameUTFLength {
			isNameCheck = false
		}
	}
	return isNameCheck
}
