package user

import (
	"github.com/rudianto-dev/gotemp-sdk/pkg/transporter"
)

type Service struct {
	tp transporter.IHttpTransporter
}

func NewService(tp transporter.IHttpTransporter) IUserService {
	return &Service{
		tp: tp,
	}
}
