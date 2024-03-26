package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

func (s *Service) GetDetail(ctx context.Context, req userContract.GetDetailUserRequest) (res *userContract.GetDetailUserResponse, err error) {
	route := fmt.Sprintf("internal/v1/user/%s", req.ID)
	execute, err := s.tp.Execute(ctx, http.MethodGet, route, req)
	if err != nil {
		return
	}
	var contentByte []byte
	contentByte, err = json.Marshal(&execute.Content)
	if err != nil {
		return
	}
	err = json.Unmarshal(contentByte, &res)
	return
}

func (s *Service) GetList(ctx context.Context, req userContract.GetListUserRequest) (res *userContract.GetListUserResponse, err error) {
	return &userContract.GetListUserResponse{}, nil
}

func (s *Service) CreateUser(ctx context.Context, req userContract.CreateUserRequest) (res *userContract.CreateUserResponse, err error) {
	return &userContract.CreateUserResponse{}, nil
}

func (s *Service) UpdateUser(ctx context.Context, req userContract.UpdateUserRequest) (res *userContract.UpdateUserResponse, err error) {
	return &userContract.UpdateUserResponse{}, nil
}

func (s *Service) DeleteUser(ctx context.Context, req userContract.DeleteUserRequest) (res *userContract.DeleteUserResponse, err error) {
	return &userContract.DeleteUserResponse{}, nil
}
