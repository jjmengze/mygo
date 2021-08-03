package errors

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"net/http"
	"strings"
)

type _error struct {
	Status int `json:"status"`
	//Code     string                 `json:"code"`
	GRPCCode codes.Code `json:"grpccode"`
	Message  string     `json:"message"`
	//Details  map[string]interface{} `json:"details"`
}

func NewErrors(status int, grpcCode codes.Code, message string) *_error {
	return &_error{
		Status:   status,
		GRPCCode: grpcCode,
		Message:  message,
	}
}

var (
	ErrInternalError = NewErrors(http.StatusInternalServerError, codes.Internal, "The server encountered an internal error. Please retry the request.")
)

func (e *_error) Error() string {
	var b strings.Builder
	if e.Status != 0 {
		b.WriteString(fmt.Sprintf("[http code]:%v,", e.Status))
	}
	b.WriteString(fmt.Sprintf("[GRPC code]:%v,", e.GRPCCode))
	b.WriteString(e.Message)
	return b.String()
}
