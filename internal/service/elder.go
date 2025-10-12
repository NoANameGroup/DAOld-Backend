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
	DeleteMyElder(ctx context.Context) (*elder.DeleteMyElderResp, error)
	UpdateMyElder(ctx context.Context, req *elder.UpdateMyElderReq) (*elder.UpdateMyElderResp, error)
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

/*
对业务逻辑的理解：DeleteMyAccount 已经存在于 UserService 中，
所以DeleteMyElder是 CreateElder 的逆操作；
更新我的Elder信息，唯一合理且安全的、可由用户直接更新的字段是 BlockChainAddress
*/

func (s *ElderService) DeleteMyElder(ctx context.Context) (*elder.DeleteMyElderResp, error) {
	var err error

	// 获取当前用户ID并转换类型
	userId, ok := ctx.Value(consts.ContextUserID).(bson.ObjectID)
	if !ok {
		return nil, errorx.ErrContextUserIDInvalid
	}

	// 将角色还原为普通用户
	if err = s.UserRepository.UpdateUserRoleByUserID(ctx, userId, enum.RoleUser); err != nil {
		log.CtxError(ctx, "failed to update user role back to user: %v", err)
		// 即使此步失败，也应该继续尝试删除elder记录，以进行数据清理，但在有事务的情况下，整个操作会直接回滚
		return nil, err
	}

	// 从elders集合删除对应记录
	if err = s.ElderRepository.DeleteElderByUserID(ctx, userId); err != nil {
		log.CtxError(ctx, "failed to delete elder record: %v", err)
		// 此处可能导致数据不一致。如果角色更新成功但删除失败，用户角色是User，但elder记录依然存在。这就是为什么需要事务。
		return nil, err
	}

	// 更新用户的 UpdatedAt 时间戳
	if err = s.UserRepository.UpdateUpdatedAtByUserID(ctx, userId, time.Now()); err != nil {
		log.CtxError(ctx, "failed to update user updated at: %v", err)
		return nil, err
	}

	return &elder.DeleteMyElderResp{Resp: dto.Success()}, nil
}

func (s *ElderService) UpdateMyElder(ctx context.Context, req *elder.UpdateMyElderReq) (*elder.UpdateMyElderResp, error) {
	// 获取当前用户ID并转换类型
	userId, ok := ctx.Value(consts.ContextUserID).(bson.ObjectID)
	if !ok {
		return nil, errorx.ErrContextUserIDInvalid
	}

	update := bson.M{}
	updateCount := 0

	// 检查请求中是否有提供BlockChainAddress字段并进行更新
	if req.BlockChainAddress != "" {
		update["blockChainAddress"] = req.BlockChainAddress
		updateCount++
	}

	// 如果没有更新的字段，直接返回成功
	if updateCount == 0 {
		return &elder.UpdateMyElderResp{Resp: dto.Success()}, nil
	}

	update[consts.UpdatedAt] = time.Now()
	if err := s.ElderRepository.UpdateElderByUserID(ctx, userId, update); err != nil {
		log.CtxError(ctx, "failed to update elder: %v", err)
		return nil, err
	}

	return &elder.UpdateMyElderResp{Resp: dto.Success()}, nil
}
