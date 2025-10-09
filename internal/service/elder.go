package service

import (
	"context"

	"github.com/NoANameGroup/DAOld-Backend/internal/dto"
	"github.com/NoANameGroup/DAOld-Backend/internal/dto/elder"
	"github.com/NoANameGroup/DAOld-Backend/internal/model"
	"github.com/NoANameGroup/DAOld-Backend/internal/repository"
	"github.com/NoANameGroup/DAOld-Backend/pkg/consts"
	"github.com/NoANameGroup/DAOld-Backend/pkg/consts/enum"
	"github.com/NoANameGroup/DAOld-Backend/pkg/errorx"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ IElderService = (*ElderService)(nil)

type IElderService interface {
	CreateElder(ctx context.Context) (*elder.CreateElderResp, error)
}

type ElderService struct {
	ElderRepository *repository.ElderRepository
	UserRepository  *repository.UserRepository
}

var ElderServiceSet = wire.NewSet(
	wire.Struct(new(ElderService), "*"),
	wire.Bind(new(IElderService), new(*ElderService)),
)

func (s *ElderService) CreateElder(ctx context.Context) (*elder.CreateElderResp, error) {
	var err error

	// 获取用户ID并转换类型
	userId, ok := ctx.Value(consts.ContextUserID).(primitive.ObjectID)
	if !ok {
		return nil, errorx.ErrContextUserIDInvalid
	}

	// 修改用户角色为老人
	if err = s.UserRepository.UpdateUserRole(ctx, userId, enum.RuleElder); err != nil {
		log.CtxError(ctx, "failed to update user role: %v", err)
		return nil, err
	}

	// 创建老人
	newElder := &model.Elder{
		ID:      primitive.NewObjectID(),
		UserID:  userId,
		Balance: 0,
	}

	// 插入数据库
	if err = s.ElderRepository.Insert(ctx, newElder); err != nil {
		log.CtxError(ctx, "failed to insert elder: %v", err)
		return nil, err
	}

	return &elder.CreateElderResp{Resp: dto.Success()}, nil
}
