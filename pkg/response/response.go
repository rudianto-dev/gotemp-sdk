package response

import (
	"fmt"
	"net/http"

	// "github.com/opentracing/opentracing-go"
	// "github.com/opentracing/opentracing-go/ext"
	// "github.com/opentracing/opentracing-go/log"
	// "github.com/roserocket/xerrs"

	// "gitlab.appsnp.ocbcnisp.com/onelabs/auth/pkg/validator"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
	// validation "github.com/go-ozzo/ozzo-validation"
)

const ErrMaxStack = 5

type (
	Causer interface {
		Cause() error
	}
	Response struct {
		RequestId string      `json:"request_id"`
		Content   interface{} `json:"content"`
		Error     *Error      `json:"error,omitempty"`
		Status    int         `json:"status"`
		Message   string      `json:"message"`
	}
	Error struct {
		Code    utils.BusinessCode `json:"code"`
		Message string             `json:"message"`
		Reasons map[string]string  `json:"reasons"`
		Details []interface{}      `json:"details,omitempty"`
	}
)

func (err *Error) Error() string {
	return fmt.Sprintf("error with code: %d; message: %s", err.Code, err.Message)
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Yay(w http.ResponseWriter, r *http.Request, status int, content interface{}) {
	// routePath := r.URL.Path
	// if r.URL.RawPath != "" {
	// 	routePath = r.URL.RawPath
	// }

	// span, _ := opentracing.StartSpanFromContext(r.Context(), routePath)
	// defer span.Finish()

	// span.SetTag("path", routePath)
	// span.SetTag("status_code", status)
	// span.SetTag("response", content)
	// span.SetTag("request_id", middleware.GetReqID(r.Context()))

	render.Status(r, status)
	_ = render.Render(w, r, &Response{
		RequestId: middleware.GetReqID(r.Context()),
		Content:   content,
		Status:    status,
		Message:   http.StatusText(status),
	})
}

func DataYay(w http.ResponseWriter, r *http.Request, filename string, content []byte) {
	contentDisposition := fmt.Sprintf("attachment;filename=%s", filename)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", contentDisposition)
	w.Write(content)
}

func Nay(w http.ResponseWriter, r *http.Request, err error, details ...interface{}) {
	var (
		message = err.Error()
	)
	// routePath := r.URL.Path
	// if r.URL.RawPath != "" {
	// 	routePath = r.URL.RawPath
	// }

	// span, _ := opentracing.StartSpanFromContext(r.Context(), routePath)
	// defer span.Finish()

	// maskErr := xerrs.Details(err, ErrMaxStack)
	// span.SetTag("path", routePath)
	// span.SetTag("status_code", status)
	// span.SetTag("error", xerrs.Details(err, ErrMaxStack))
	// span.SetTag("request_id", middleware.GetReqID(r.Context()))
	// ext.LogError(span, errors.New(maskErr), log.String("code", fmt.Sprintf("%d", status)))

	// map error to http status header
	httpStatus := http.StatusInternalServerError
	businessCode := utils.Undefine
	// handle validator error catch
	reasons := utils.GetValidatorFieldError(err)
	if len(reasons) > 0 {
		err = utils.ErrBadRequest
	}
	// handle business error catch
	if val, ok := utils.BusinessToError[err]; ok {
		bm := utils.GetStatusMessage(val)
		businessCode = val
		message = bm.Message
		httpStatus = bm.HttpStatusCode
	}
	render.Status(r, httpStatus)
	_ = render.Render(w, r, &Response{
		RequestId: middleware.GetReqID(r.Context()),
		Status:    httpStatus,
		Message:   http.StatusText(httpStatus),
		Error: &Error{
			Code:    businessCode,
			Message: message,
			Reasons: reasons,
			Details: details,
		},
	})
}
