package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

func (s *Service) GetDetail(ctx context.Context, req userContract.GetDetailUserRequest) (res *userContract.GetDetailUserResponse, err error) {
	route := fmt.Sprintf("/v1/user/%s", req.ID)
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

func (s *Service) GetList(ctx context.Context, req userContract.GetListUserRequest) (*userContract.GetListUserResponse, error) {
	return nil, nil
}

func (s *Service) CreateUser(ctx context.Context, req userContract.CreateUserRequest) (*userContract.CreateUserResponse, error) {
	return nil, nil
}

func (s *Service) UpdateUser(ctx context.Context, req userContract.UpdateUserRequest) (*userContract.UpdateUserResponse, error) {
	return nil, nil
}

func (s *Service) DeleteUser(ctx context.Context, req userContract.DeleteUserRequest) (*userContract.DeleteUserResponse, error) {
	return nil, nil
}
