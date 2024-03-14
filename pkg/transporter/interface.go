package transporter

import "context"

type (
	ErrorResponse struct {
		Code    int               `json:"code"`
		Message string            `json:"message"`
		Reasons map[string]string `json:"reasons"`
	}
	Response struct {
		RequestID string                 `json:"request_id"`
		Content   map[string]interface{} `json:"content,omitempty"`
		Status    int                    `json:"status"`
		Message   string                 `json:"message"`
		Error     *ErrorResponse         `json:"error,omitempty"`
	}
)

type IHttpTransporter interface {
	Execute(ctx context.Context, method, route string, payload interface{}) (resp *Response, err error)
	ExecuteWithToken(ctx context.Context, method, route, token string, payload interface{}) (resp *Response, err error)
}
