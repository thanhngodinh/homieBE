package domain

var USER_SECRET_KEY = []byte("user_key")

type LoginReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginRes struct {
	IsResetPass bool         `json:"isResetPass,omitempty"`
	Token       string       `json:"token,omitempty"`
	Profile     *UserProfile `json:"profile,omitempty"`
}

type UpdatePasswordReq struct {
	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}

type RegisterReq struct {
	Username string `json:"username,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Name     string `json:"name,omitempty"`
}

type VerifyPhoneReq struct {
	Phone string `json:"phone,omitempty"`
}

type VerifyOTPReq struct {
	OTP string `json:"otp,omitempty"`
}
