package resterr

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"err"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// variavel Error do Go eh uma interface entao vc pode adicionar coisas nela

func (r *RestErr) Error() string {
	return r.Message
}

func NewBadRequesterror(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Causes:  nil,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Causes:  nil,
	}
}
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Causes:  nil,
	}
}
