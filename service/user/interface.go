package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
	utilContract "github.com/rudianto-dev/gotemp-sdk/contract/util"
)

type IUserService interface {
	Health(ctx context.Context) (res *utilContract.GetHealthResponse, err error)

	GetDetail(ctx context.Context, req userContract.GetDetailUserRequest) (res *userContract.GetDetailUserResponse, err error)
	GetList(ctx context.Context, req userContract.GetListUserRequest) (res *userContract.GetListUserResponse, err error)
	CreateUser(ctx context.Context, req userContract.CreateUserRequest) (res *userContract.CreateUserResponse, err error)
	UpdateUser(ctx context.Context, req userContract.UpdateUserRequest) (res *userContract.UpdateUserResponse, err error)
	DeleteUser(ctx context.Context, req userContract.DeleteUserRequest) (res *userContract.DeleteUserResponse, err error)
}
