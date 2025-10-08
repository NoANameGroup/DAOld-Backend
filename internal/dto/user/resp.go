package user

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/dto"
)

type CreateUserResp struct {
	*dto.Resp
}

type GetMyProfileResp struct {
	*dto.Resp
	*UserVO
}

type UpdateMyPasswordResp struct {
	*dto.Resp
}

type DeleteMyAccountResp struct {
	*dto.Resp
}

type UpdateMyProfileResp struct {
	*dto.Resp
	Count int `json:"count"`
}

type UpdateUserRoleResp struct {
	*dto.Resp
}

type UpdateMyEmailResp struct {
	*dto.Resp
}

type UpdateMyPhoneResp struct {
	*dto.Resp
}

type UpdateMyAvatarResp struct {
	*dto.Resp
}

type UpdateMyUsernameResp struct {
	*dto.Resp
}
