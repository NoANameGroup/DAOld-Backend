package user

import "github.com/NoANameGroup/DAOld-Backend/adaptor/dto"

type RegisterResp struct {
	*dto.Resp
}

type LoginResp struct {
	*dto.Resp
	AccessToken string `json:"accessToken"`
}
