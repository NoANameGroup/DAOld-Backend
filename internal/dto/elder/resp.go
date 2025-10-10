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

type DeleteMyElderResp struct {
	*dto.Resp
}
