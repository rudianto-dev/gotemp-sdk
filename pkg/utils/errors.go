package utils

import (
	"errors"
	"net/http"
)

type Causer interface {
	Cause() error
}

type ErrorWithCode struct {
	code  int
	cause error
}

func NewErrorWithCode(c int, e error) *ErrorWithCode {
	return &ErrorWithCode{code: c, cause: e}
}

func (err *ErrorWithCode) Error() string {
	return err.Cause().Error()
}

func (err *ErrorWithCode) Code() int {
	return err.code
}

func (err *ErrorWithCode) Cause() error {
	return err.cause
}

type ErrorStage string

var (
	ERROR_REPOSITORY_STAGE ErrorStage = "repository"
	ERROR_USE_CASE_STAGE   ErrorStage = "usecase"
	ERROR_HANDLER_STAGE    ErrorStage = "handler"
	ERROR_INFRA_STAGE      ErrorStage = "infra"
)

type ErrorType error

var (
	ErrNotFound      ErrorType = errors.New("data not found")
	ErrForbidden     ErrorType = errors.New("forbidden access")
	ErrBadRequest    ErrorType = errors.New("bad request")
	ErrQueryRead     ErrorType = errors.New("error while querying storage")
	ErrQueryScan     ErrorType = errors.New("error while scan query result")
	ErrQueryTxBegin  ErrorType = errors.New("error while preparing query transaction")
	ErrQueryTxInsert ErrorType = errors.New("error while inserting record")
	ErrQueryTxUpdate ErrorType = errors.New("error while updating record")
	ErrQueryTxDelete ErrorType = errors.New("error while deleting record")
	ErrQueryTxCommit ErrorType = errors.New("error while committing record changes")

	ErrInvalidContentType = errors.New("invalid content type")
)

type CustomErrorMap struct {
	BusinessErrorCode int
	HttpStatusCode    int
	Origin            error
}

var (
	ErrorMapper = map[ErrorType]CustomErrorMap{
		ErrNotFound:           {BusinessErrorCode: 101, HttpStatusCode: http.StatusNotFound, Origin: ErrNotFound},
		ErrForbidden:          {BusinessErrorCode: 102, HttpStatusCode: http.StatusForbidden, Origin: ErrForbidden},
		ErrBadRequest:         {BusinessErrorCode: 103, HttpStatusCode: http.StatusBadRequest, Origin: ErrBadRequest},
		ErrQueryRead:          {BusinessErrorCode: 104, HttpStatusCode: http.StatusInternalServerError, Origin: ErrQueryRead},
		ErrQueryScan:          {BusinessErrorCode: 105, HttpStatusCode: http.StatusInternalServerError, Origin: ErrQueryScan},
		ErrQueryTxBegin:       {BusinessErrorCode: 106, HttpStatusCode: http.StatusInternalServerError, Origin: ErrQueryTxBegin},
		ErrQueryTxInsert:      {BusinessErrorCode: 107, HttpStatusCode: http.StatusInternalServerError, Origin: ErrQueryTxInsert},
		ErrQueryTxUpdate:      {BusinessErrorCode: 108, HttpStatusCode: http.StatusInternalServerError, Origin: ErrQueryTxUpdate},
		ErrQueryTxDelete:      {BusinessErrorCode: 109, HttpStatusCode: http.StatusInternalServerError, Origin: ErrQueryTxDelete},
		ErrQueryTxCommit:      {BusinessErrorCode: 110, HttpStatusCode: http.StatusInternalServerError, Origin: ErrQueryTxCommit},
		ErrInvalidContentType: {BusinessErrorCode: 111, HttpStatusCode: http.StatusBadRequest, Origin: ErrInvalidContentType},
	}
)
