package elder

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/dto"
)

type CreateElderResp struct {
	*dto.Resp
}

type GetMyElderResp struct {
	*dto.Resp
	*ElderVO
}

type UpdateMyElderResp struct {
	*dto.Resp
	Count int `json:"count"`
}

type DeleteMyElderResp struct {
	*dto.Resp
}
