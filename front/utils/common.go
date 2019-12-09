package utils

import (
	"regexp"
)

func EmailCheck(email string) bool {
	var rxEmail = regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)
	isEmail := rxEmail.MatchString(email)
	return isEmail
}
