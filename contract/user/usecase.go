package user

type UserRequest struct {
	Name string `json:"name" validate:"required"`
}
type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
}

type GetProfileRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetProfileResponse struct {
	User UserResponse `json:"user"`
}

type GetListRequest struct {
}

type GetListResponse struct {
	Users []UserResponse `json:"users"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateUserResponse struct {
	User UserResponse `json:"user"`
}

type UpdateUserRequest struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateUserResponse struct {
	User UserResponse `json:"user"`
}

type DeleteUserRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeleteUserResponse struct {
	User UserResponse `json:"user"`
}
