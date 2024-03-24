package token

type Payload struct {
	ID       string `json:"id"`
	RoleType int8   `json:"role_type"`
}
