package domain

var ADMIN_SECRET_KEY = []byte("admin_key")

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token   string   `json:"token"`
	Profile *Profile `json:"profile"`
}

type Profile struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar,omitempty"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
