package token

type Payload struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	RoleType int8   `json:"role_type"`
}