package user

import "github.com/rudianto-dev/gotemp-sdk/pkg/transporter"

type Gateway struct {
	tp transporter.HTTPTransporter
}

func NewHandler(tp transporter.HTTPTransporter) IUserGateway {
	return &Gateway{
		tp: tp,
	}
}
