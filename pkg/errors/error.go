package errors

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
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
	ErrInternalError         = NewErrors(http.StatusInternalServerError, codes.Internal, "The server encountered an internal error. Please retry the request.")
	ErrResourceNotFound      = NewErrors(http.StatusNotFound, codes.NotFound, "The specified resource does not exist.")
	ErrResourceAlreadyExists = NewErrors(http.StatusConflict, codes.AlreadyExists, "The specified resource already exists.")
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

var (
	ErrInvalidInput = &_error{Message: "One of the request inputs is not valid.", Status: http.StatusBadRequest, GRPCCode: codes.InvalidArgument}
)

// ConvertMySQLError convert mysql error
func ConvertMySQLError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrResourceNotFound
	}

	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		if mysqlErr.Number == 1062 {
			// the duplicate key error.
			return ErrResourceAlreadyExists
		}
	}

	return ErrInternalError
}
