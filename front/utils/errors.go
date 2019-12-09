package utils

type ErrorMessage struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
	ErrorCode  error  `json:"error_code`
}

func (ErrorMessage) ErrorMsg(msg string, code int, errorCode error) *ErrorMessage {
	return &ErrorMessage{
		Message:    msg,
		StatusCode: code,
		ErrorCode:  errorCode,
	}
}
