package utils

import (
	"regexp"
	"unicode/utf8"
)

const (
	passwordLength = 8 //비밀번호 최소 길이
	nameUTFLength  = 2 //한글 이름 최소길이
	utfFusionLenth = 3 //한글, 영문 조합 최소길이
	nameLength     = 4 //영문이릉 최소길이
)

func EmailCheck(email string) bool {
	var rxEmail = regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)
	isEmail := rxEmail.MatchString(email)
	return isEmail
}

//비밀번호은 영문조합 8자리 이상
func PasswordCheck(password string) bool {

	pattern := "[[:graph:]]{8}"
	isPassword, _ := regexp.Match(pattern, []byte(password))
	return isPassword
}

//이름체크
func NameCheck(name string) (bool, int) {
	utfCount := utf8.RuneCountInString(name) //문자열의 실제 길이를 구함.
	normalCount := len(name)

	isNameCheck := true
	errorCount := 0

	if normalCount == utfCount {
		//영문이름일 경우 4글자 이상
		if normalCount < nameLength {
			isNameCheck = false
			errorCount = nameLength
		}
	} else {
		//한글이름일 경우 2글자 이상
		if utfCount < nameUTFLength {
			isNameCheck = false
			errorCount = nameUTFLength
		} else if normalCount == nameLength {
			isNameCheck = false
			errorCount = utfFusionLenth
		}
	}
	return isNameCheck, errorCount
}
