package domain

var ADMIN_SECRET_KEY = []byte("admin_key")

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token   string   `json:"token,omitempty"`
	Profile *Profile `json:"profile,omitempty"`
}

type Profile struct {
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}
