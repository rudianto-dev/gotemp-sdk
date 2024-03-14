package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

type IUserService interface {
	GetDetail(ctx context.Context, req userContract.GetDetailUserRequest) (*userContract.GetDetailUserResponse, error)
	GetList(ctx context.Context, req userContract.GetListUserRequest) (*userContract.GetListUserResponse, error)
	CreateUser(ctx context.Context, req userContract.CreateUserRequest) (*userContract.CreateUserResponse, error)
	UpdateUser(ctx context.Context, req userContract.UpdateUserRequest) (*userContract.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req userContract.DeleteUserRequest) (*userContract.DeleteUserResponse, error)
}
