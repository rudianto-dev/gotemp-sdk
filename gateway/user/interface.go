package user

import (
	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

type IUserGateway interface {
	GetUserProfile(req userContract.GetProfileRequest) (*userContract.GetProfileResponse, error)
}
