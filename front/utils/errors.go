package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
}

func (ErrorMessage) ReturnErrorMsg(msg string, code int) *ErrorMessage {
	return &ErrorMessage{
		Message:    msg,
		StatusCode: code,
	}
}

func (ErrorMessage) ErrorMsg(w http.ResponseWriter, msg string, code int) {
	message := ErrorMessage{
		Message:    msg,
		StatusCode: code,
	}
	w.Header().Set("Content-Type", "application/json")
	outData, _ := json.Marshal(message)
	w.WriteHeader(code)
	w.Write(outData)
}
