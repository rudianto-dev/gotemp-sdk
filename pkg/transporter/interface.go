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
		Status    int                    `json:"status"`
		Content   map[string]interface{} `json:"content,omitempty"`
		Error     *ErrorResponse         `json:"error,omitempty"`
	}
)

type IHttpTransporter interface {
	Execute(ctx context.Context, method, url, token string, payload interface{}) (resp *Response, err error)
	ExecuteWithToken(ctx context.Context, method, url, token string, payload interface{}) (resp *Response, err error)
}
