package user

type RegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordReq struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type DeleteAccountReq struct {
	Password     string `json:"password"`
	Confirmation string `json:"confirmation"`
}

type UpdateMyProfileReq struct {
	*UserVO
}

type UpdateUserRoleReq struct {
	Role string `json:"role"`
}
