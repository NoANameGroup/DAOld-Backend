package service

import (
	"context"
	"time"

	"github.com/NoANameGroup/DAOld-Backend/internal/dto"
	"github.com/NoANameGroup/DAOld-Backend/internal/dto/elder"
	"github.com/NoANameGroup/DAOld-Backend/internal/model"
	"github.com/NoANameGroup/DAOld-Backend/internal/repository"
	"github.com/NoANameGroup/DAOld-Backend/pkg/consts"
	"github.com/NoANameGroup/DAOld-Backend/pkg/consts/enum"
	"github.com/NoANameGroup/DAOld-Backend/pkg/errorx"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var _ IElderService = (*ElderService)(nil)

type IElderService interface {
	CreateElder(ctx context.Context) (*elder.CreateElderResp, error)
	GetMyElder(ctx context.Context) (*elder.GetMyElderResp, error)
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
	userId, ok := ctx.Value(consts.ContextUserID).(bson.ObjectID)
	if !ok {
		return nil, errorx.ErrContextUserIDInvalid
	}

	// 修改用户角色为老人
	if err = s.UserRepository.UpdateUserRoleByUserID(ctx, userId, enum.RoleElder); err != nil {
		log.CtxError(ctx, "failed to update user role: %v", err)
		return nil, err
	}

	// 修改 UpdateAt 为当前时间
	if err = s.UserRepository.UpdateUpdatedAtByUserID(ctx, userId, time.Now()); err != nil {
		log.CtxError(ctx, "failed to update updated at: %v", err)
		return nil, err
	}

	// 创建老人
	newElder := &model.Elder{
		ID:        bson.NewObjectID(),
		UserID:    userId,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 插入数据库
	if err = s.ElderRepository.Insert(ctx, newElder); err != nil {
		log.CtxError(ctx, "failed to insert elder: %v", err)
		return nil, err
	}

	return &elder.CreateElderResp{Resp: dto.Success()}, nil
}

func (s *ElderService) GetMyElder(ctx context.Context) (*elder.GetMyElderResp, error) {
	var err error
	var userModel *model.User
	var elderModel *model.Elder

	// 获取用户ID并转换类型
	userId, ok := ctx.Value(consts.ContextUserID).(bson.ObjectID)
	if !ok {
		return nil, errorx.ErrContextUserIDInvalid
	}

	// 获取老人信息
	if elderModel, err = s.ElderRepository.FindElderByUserID(ctx, userId); err != nil {
		log.CtxError(ctx, "failed to find elder: %v", err)
		return nil, err
	}

	// 获取用户信息
	if userModel, err = s.UserRepository.FindUserByUserID(ctx, userId); err != nil {
		log.CtxInfo(ctx, "failed to find user: %v", err)
		return nil, err
	}

	log.CtxInfo(ctx, "userModel: %v", userModel)

	return &elder.GetMyElderResp{
		Resp: dto.Success(),
		ElderVO: &elder.ElderVO{
			Username:          userModel.Username,
			BlockChainAddress: elderModel.BlockChainAddress,
			Balance:           elderModel.Balance,
			CreatedAt:         elderModel.CreatedAt,
		},
	}, nil
}
