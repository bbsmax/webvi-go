package utils

import (
	"encoding/json"
	"net/http"
)

type ReturnMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type ReturnData struct {
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
	Status string      `json:"status"`
}

func (ReturnMessage) ReturnMsg(msg string, code int, status string) *ReturnMessage {
	return &ReturnMessage{
		Message: msg,
		Code:    code,
		Status:  status,
	}
}

func (ReturnMessage) ResponseMsg(w http.ResponseWriter, msg string, code int, status string) {
	message := ReturnMessage{
		Message: msg,
		Code:    code,
		Status:  status,
	}
	w.Header().Set("Content-Type", "application/json")
	outData, _ := json.Marshal(message)
	w.WriteHeader(code)
	w.Write(outData)
}

func (ReturnMessage) ResponseData(w http.ResponseWriter, msg interface{}, code int, status string) {
	message := ReturnData{
		Data:   msg,
		Code:   code,
		Status: status,
	}
	w.Header().Set("Content-Type", "application/json")
	outData, _ := json.Marshal(message)
	w.WriteHeader(code)
	w.Write(outData)
}
