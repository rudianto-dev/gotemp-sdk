package utils

import "net/http"

type BusinessCode int

const (
	Success                   BusinessCode = 200
	SuccessCreated            BusinessCode = 201
	SuccessNoContent          BusinessCode = 204
	BadRequest                BusinessCode = 400
	Unauthorized              BusinessCode = 401
	Forbidden                 BusinessCode = 403
	NotFound                  BusinessCode = 404
	MethodNotAllowed          BusinessCode = 405
	RequestTimeout            BusinessCode = 408
	TooManyRequests           BusinessCode = 429
	QueryReadInvalid          BusinessCode = 460
	QueryScanInvalid          BusinessCode = 461
	QueryTxBeginInvalid       BusinessCode = 462
	QueryTxInsertInvalid      BusinessCode = 463
	QueryTxUpdateInvalid      BusinessCode = 464
	QueryTxDeleteInvalid      BusinessCode = 465
	QueryTxCommitInvalid      BusinessCode = 466
	ContentTypeInvalid        BusinessCode = 467
	UsernameAlreadyRegistered BusinessCode = 468
	InvalidCredential         BusinessCode = 469
	ExpireOTP                 BusinessCode = 470
	ExpireOTPVerification     BusinessCode = 471
	Undefine                  BusinessCode = 499
	InternalError             BusinessCode = 500
	BadGateway                BusinessCode = 502
	ServiceUnavailable        BusinessCode = 503
)

type BusinessStatusCode struct {
	Message        string
	HttpStatusCode int
}

var BusinessStatusMessage = map[BusinessCode]BusinessStatusCode{
	Success:                   {Message: "request success", HttpStatusCode: http.StatusOK},
	SuccessCreated:            {Message: "created success", HttpStatusCode: http.StatusCreated},
	SuccessNoContent:          {Message: "success with no content", HttpStatusCode: http.StatusNoContent},
	BadRequest:                {Message: "bad request", HttpStatusCode: http.StatusBadRequest},
	Unauthorized:              {Message: "unauthorized", HttpStatusCode: http.StatusUnauthorized},
	Forbidden:                 {Message: "forbidden", HttpStatusCode: http.StatusForbidden},
	NotFound:                  {Message: "data not found", HttpStatusCode: http.StatusNotFound},
	MethodNotAllowed:          {Message: "request not allowed", HttpStatusCode: http.StatusMethodNotAllowed},
	RequestTimeout:            {Message: "request time out", HttpStatusCode: http.StatusRequestTimeout},
	TooManyRequests:           {Message: "to many request", HttpStatusCode: http.StatusTooManyRequests},
	QueryReadInvalid:          {Message: "error while querying storage", HttpStatusCode: http.StatusInternalServerError},
	QueryScanInvalid:          {Message: "error when scan query result", HttpStatusCode: http.StatusInternalServerError},
	QueryTxBeginInvalid:       {Message: "error while preparing query transaction", HttpStatusCode: http.StatusInternalServerError},
	QueryTxInsertInvalid:      {Message: "error while inserting record", HttpStatusCode: http.StatusInternalServerError},
	QueryTxUpdateInvalid:      {Message: "error while updating record", HttpStatusCode: http.StatusInternalServerError},
	QueryTxDeleteInvalid:      {Message: "error while deleting record", HttpStatusCode: http.StatusInternalServerError},
	QueryTxCommitInvalid:      {Message: "error while committing record changes", HttpStatusCode: http.StatusInternalServerError},
	ContentTypeInvalid:        {Message: "invalid content type", HttpStatusCode: http.StatusInternalServerError},
	UsernameAlreadyRegistered: {Message: "Username already registered", HttpStatusCode: http.StatusBadRequest},
	Undefine:                  {Message: "something went wrong", HttpStatusCode: http.StatusInternalServerError},
	InternalError:             {Message: "internal server error", HttpStatusCode: http.StatusInternalServerError},
	BadGateway:                {Message: "bad gateway", HttpStatusCode: http.StatusBadGateway},
	InvalidCredential:         {Message: "wrong username or password", HttpStatusCode: http.StatusBadRequest},
	ExpireOTP:                 {Message: "otp not valid or has been expired", HttpStatusCode: http.StatusBadRequest},
	ExpireOTPVerification:     {Message: "otp verification not valid or has been expired", HttpStatusCode: http.StatusBadRequest},
	ServiceUnavailable:        {Message: "service unavailable", HttpStatusCode: http.StatusServiceUnavailable},
}

func GetStatusMessage(code BusinessCode) BusinessStatusCode {
	if val, ok := BusinessStatusMessage[code]; ok {
		return val
	}
	return BusinessStatusMessage[Undefine]
}
