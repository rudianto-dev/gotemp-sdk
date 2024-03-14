package transporter

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type HTTPTransporter struct {
	hc       *http.Client
	url, uid string
}

const (
	mimeApplicationJson = "application/json"
)

func NewHTTPTransporter(client *http.Client, uid, url string) IHttpTransporter {
	return &HTTPTransporter{client, uid, url}
}

func (tp HTTPTransporter) Execute(ctx context.Context, method, url, token string, payload interface{}) (resp *Response, err error) {
	resp, err = tp.talk(method, url, "", payload)
	if err != nil {
		return
	}
	return
}

func (tp HTTPTransporter) ExecuteWithToken(ctx context.Context, method, url, token string, payload interface{}) (resp *Response, err error) {
	resp, err = tp.talk(method, url, "", payload)
	if err != nil {
		return
	}
	return
}

func (tp *HTTPTransporter) talk(method, url, token string, payload interface{}) (*Response, error) {
	var ir io.Reader
	if nil != payload {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, errors.WithMessage(err, "failed encoding request payload")
		}
		ir = bytes.NewReader(b)
	}
	r, err := tp.build(method, url, token, ir)
	if err != nil {
		return nil, err
	}
	return tp.request(r)
}

func (tp *HTTPTransporter) build(method, url, token string, payload io.Reader) (*http.Request, error) {
	r, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, errors.WithMessage(err, "failed creating request")
	}

	if token != "" {
		// change to bearer
		r.Header.Set("X-Auth-Token", token)
	}
	r.Header.Set("Content-Type", mimeApplicationJson)
	r.Header.Set("Accept", mimeApplicationJson)
	r.Header.Set("X-Device-Id", tp.uid)
	r.Header.Set("Cache-Control", "no-cache")
	return r, nil
}

func (tp *HTTPTransporter) request(r *http.Request) (*Response, error) {

	resp, err := tp.hc.Do(r)
	if err != nil {
		return nil, errors.WithMessage(err, "failed communicating with upstream")
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusUnauthorized {
		return &Response{Status: resp.StatusCode}, errors.New(http.StatusText(resp.StatusCode))
	}

	var rs Response
	if err := json.NewDecoder(resp.Body).Decode(&rs); err != nil {
		return nil, errors.WithMessage(err, "failed decoding response")
	}

	return &rs, nil
}
