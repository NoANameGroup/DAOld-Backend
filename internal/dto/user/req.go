package user

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateMyPasswordReq struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type DeleteMyAccountReq struct {
	Password     string `json:"password"`
	Confirmation string `json:"confirmation"`
}

type UpdateMyProfileReq struct {
	*UserVO
}

type UpdateUserRoleReq struct {
	Role string `json:"role"`
}

type UpdateMyEmailReq struct {
	Password         string `json:"password"`
	CurrentEmailCode string `json:"currentEmailCode"`
	NewEmail         string `json:"newEmail"`
	NewEmailCode     string `json:"newEmailCode"`
}

type UpdateMyPhoneReq struct {
	Password         string `json:"password"`
	CurrentPhoneCode string `json:"currentPhoneCode"`
	NewPhone         string `json:"newPhone"`
	NewPhoneCode     string `json:"newPhoneCode"`
}

type UpdateMyUsernameReq struct {
	NewUsername string `json:"newUsername"`
}
