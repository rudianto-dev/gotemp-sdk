package response

type (
	ResponseSuccess struct {
		RequestId string      `json:"request_id" example:"onelabs/IkhB54kdJd-000006"`
		Content   interface{} `json:"content"`
		Status    int         `json:"status" example:"200"`
	}

	ResponseBadRequest struct {
		RequestId string           `json:"request_id" example:"onelabs/IkhB54kdJd-000006"`
		Error     *ErrorBadRequest `json:"error,omitempty"`
		Status    int              `json:"status" example:"400"`
	}

	ErrorBadRequest struct {
		Code    string           `json:"code" example:"400"`
		Message string           `json:"message" example:"Data is not valid"`
		Reasons map[string]error `json:"reasons"`
	}

	ResponseInternalServerError struct {
		RequestId string                    `json:"request_id" example:"onelabs/IkhB54kdJd-000006"`
		Error     *ErrorInternalServerError `json:"error,omitempty"`
		Status    int                       `json:"status" example:"500"`
	}

	ErrorInternalServerError struct {
		Code    string           `json:"code" example:"500"`
		Message string           `json:"message" example:"Something went wrong, please try again"`
		Reasons map[string]error `json:"reasons"`
	}

	ResponseForbiddenError struct {
		RequestId string               `json:"request_id" example:"onelabs/IkhB54kdJd-000006"`
		Error     *ErrorForbiddenError `json:"error,omitempty"`
		Status    int                  `json:"status" example:"403"`
	}

	ErrorForbiddenError struct {
		Code    string           `json:"code" example:"403"`
		Message string           `json:"message" example:"You don't have permission to access this"`
		Reasons map[string]error `json:"reasons"`
	}
)
