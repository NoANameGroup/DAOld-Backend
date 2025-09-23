package service

import "github.com/NoANameGroup/DAOld-Backend/adaptor/dto"

type IUserService interface {
	Register(ctx context.Context, req *dto.RegisterReq) (*dto.RegisterResp, error)
}
