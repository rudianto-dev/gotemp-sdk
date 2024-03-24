package utils

import (
	"errors"
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
	ERROR_USE_CASE_STAGE   ErrorStage = "use-case"
	ERROR_HANDLER_STAGE    ErrorStage = "handler"
	ERROR_INFRA_STAGE      ErrorStage = "infra"
)

type ErrorType error

var (
	ErrNotFound                  ErrorType = errors.New("data not found")
	ErrForbidden                 ErrorType = errors.New("forbidden access")
	ErrBadRequest                ErrorType = errors.New("bad request")
	ErrQueryRead                 ErrorType = errors.New("error while querying storage")
	ErrQueryScan                 ErrorType = errors.New("error while scan query result")
	ErrQueryTxBegin              ErrorType = errors.New("error while preparing query transaction")
	ErrQueryTxInsert             ErrorType = errors.New("error while inserting record")
	ErrQueryTxUpdate             ErrorType = errors.New("error while updating record")
	ErrQueryTxDelete             ErrorType = errors.New("error while deleting record")
	ErrQueryTxCommit             ErrorType = errors.New("error while committing record changes")
	ErrInvalidContentType        ErrorType = errors.New("invalid content type")
	ErrUsernameAlreadyRegistered ErrorType = errors.New("username already registered")
	ErrGenerateOTP               ErrorType = errors.New("error while generate otp")
	ErrExpiredOTP                ErrorType = errors.New("otp not valid or has been expired")
	ErrExpiredVerificationOTP    ErrorType = errors.New("otp verification not valid or has been expired")
	ErrInvalidOTP                ErrorType = errors.New("invalid otp code")
	ErrRepositoryOTP             ErrorType = errors.New("otp repository internal error")
	ErrInvalidCredential         ErrorType = errors.New("wrong username or password")
)

var (
	BusinessToError = map[ErrorType]BusinessCode{
		ErrNotFound:                  NotFound,
		ErrForbidden:                 Forbidden,
		ErrBadRequest:                BadRequest,
		ErrQueryRead:                 QueryReadInvalid,
		ErrQueryScan:                 QueryScanInvalid,
		ErrQueryTxBegin:              QueryTxBeginInvalid,
		ErrQueryTxInsert:             QueryTxInsertInvalid,
		ErrQueryTxUpdate:             QueryTxUpdateInvalid,
		ErrQueryTxDelete:             QueryTxDeleteInvalid,
		ErrQueryTxCommit:             QueryTxCommitInvalid,
		ErrInvalidContentType:        ContentTypeInvalid,
		ErrUsernameAlreadyRegistered: UsernameAlreadyRegistered,
		ErrInvalidCredential:         InvalidCredential,
		ErrExpiredOTP:                ExpireOTP,
		ErrExpiredVerificationOTP:    ExpireOTPVerification,
	}

	ErrorToBusiness = map[BusinessCode]ErrorType{
		NotFound:                  ErrNotFound,
		Forbidden:                 ErrForbidden,
		BadRequest:                ErrBadRequest,
		QueryReadInvalid:          ErrQueryRead,
		QueryScanInvalid:          ErrQueryScan,
		QueryTxBeginInvalid:       ErrQueryTxBegin,
		QueryTxInsertInvalid:      ErrQueryTxInsert,
		QueryTxUpdateInvalid:      ErrQueryTxUpdate,
		QueryTxDeleteInvalid:      ErrQueryTxDelete,
		QueryTxCommitInvalid:      ErrQueryTxCommit,
		UsernameAlreadyRegistered: ErrUsernameAlreadyRegistered,
		InvalidCredential:         ErrInvalidCredential,
		ExpireOTP:                 ErrExpiredOTP,
		ExpireOTPVerification:     ErrExpiredVerificationOTP,
	}
)
