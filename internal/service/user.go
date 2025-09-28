package service

import (
	"context"
	"errors"
	"github.com/NoANameGroup/DAOld-Backend/internal/consts"
	"github.com/NoANameGroup/DAOld-Backend/internal/consts/enum"
	"github.com/NoANameGroup/DAOld-Backend/internal/dto"
	"github.com/NoANameGroup/DAOld-Backend/internal/dto/user"
	"github.com/NoANameGroup/DAOld-Backend/internal/errorx"
	"github.com/NoANameGroup/DAOld-Backend/internal/jwt"
	"github.com/NoANameGroup/DAOld-Backend/internal/model"
	"github.com/NoANameGroup/DAOld-Backend/internal/repository"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/NoANameGroup/DAOld-Backend/pkg/security"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type IUserService interface {
	Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterResp, error)
	Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error)
	GetMyProfile(ctx context.Context) (*user.GetMyProfileResp, error)
	ChangePassword(ctx context.Context, req *user.ChangePasswordReq) (*user.ChangePasswordResp, error)
	DeleteAccount(ctx context.Context, req *user.DeleteAccountReq) (*user.DeleteAccountResp, error)
	UpdateMyProfile(ctx context.Context, req *user.UpdateMyProfileReq) (*user.UpdateMyProfileResp, error)
}

type UserService struct {
	UserRepository *repository.UserRepository
}

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)),
)

func (s *UserService) Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterResp, error) {
	var err error
	var hashPassword string

	// 生成哈希密码
	if hashPassword, err = security.HashPassword(req.Password); err != nil {
		log.CtxError(ctx, "failed to hash password: %v", err)
		return nil, err
	}

	// 创建用户
	newUser := &model.User{
		ID:        primitive.NewObjectID(),
		Email:     req.Email,
		Username:  req.Username,
		Password:  hashPassword,
		Role:      enum.RoleUser,
		Status:    enum.StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 插入数据库
	if err = s.UserRepository.Insert(ctx, newUser); err != nil {
		log.CtxError(ctx, "failed to insert user: %v", err)
		return nil, err
	}

	return &user.RegisterResp{Resp: dto.Success()}, nil
}

func (s *UserService) Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	var err error
	var token string
	var newUser *model.User

	// 获取用户
	if newUser, err = s.UserRepository.FindUserByEmail(ctx, req.Email); err != nil {
		log.CtxError(ctx, "failed to find user: %v", err)
		return nil, err
	}

	// 校验密码是否正确
	if !security.ComparePassword(newUser.Password, req.Password) {
		log.CtxInfo(ctx, "username or password incorrect")
		return nil, errors.New("username or password incorrect")
	}

	// 更新最后登录时间
	if err = s.UserRepository.UpdateLastLoginAt(ctx, newUser.ID, time.Now()); err != nil {
		log.CtxError(ctx, "failed to update last login at: %v", err)
		return nil, err
	}

	// 生成 token
	token, err = jwt.GenerateToken(newUser.ID)
	if err != nil {
		log.CtxError(ctx, "failed to generate token: %v", err)
		return nil, err
	}

	return &user.LoginResp{
		Resp:        dto.Success(),
		AccessToken: token,
		UserID:      newUser.ID,
	}, nil
}

func (s *UserService) GetMyProfile(ctx context.Context) (*user.GetMyProfileResp, error) {
	var err error
	var userModel *model.User

	// 获取用户ID并转换类型
	userId, ok := ctx.Value(consts.ContextUserID).(primitive.ObjectID)
	if !ok {
		return nil, errors.New("invalid user id in context")
	}

	// 获取用户信息
	userModel, err = s.UserRepository.FindUserByUserID(ctx, userId)
	if err != nil {
		log.CtxError(ctx, "failed to find user: %v", err)
		return nil, err
	}

	return &user.GetMyProfileResp{
		Resp: dto.Success(),
		UserVO: &user.UserVO{
			Email:       userModel.Email,
			Username:    userModel.Username,
			Avatar:      userModel.Avatar,
			FirstName:   userModel.FirstName,
			LastName:    userModel.LastName,
			Gender:      enum.GetUserGenderDesc(userModel.Gender),
			Role:        enum.GetUserRoleDesc(userModel.Role),
			Status:      enum.GetUserStatusDesc(userModel.Status),
			Phone:       userModel.Phone,
			Address:     userModel.Address,
			Bio:         userModel.Bio,
			Birthday:    userModel.Birthday.Format("2006-01-02"),
			LastLoginAt: userModel.LastLoginAt,
			CreatedAt:   userModel.CreatedAt,
		},
	}, nil
}

func (s *UserService) ChangePassword(ctx context.Context, req *user.ChangePasswordReq) (*user.ChangePasswordResp, error) {
	var err error
	var userModel *model.User
	var hashPassword string

	// 获取用户ID并转换类型
	userId, ok := ctx.Value(consts.ContextUserID).(primitive.ObjectID)
	if !ok {
		return nil, errors.New("invalid user id in context")
	}

	// 获取用户
	userModel, err = s.UserRepository.FindUserByUserID(ctx, userId)
	if err != nil {
		log.CtxError(ctx, "failed to find user: %v", err)
		return nil, err
	}

	// 校验旧密码是否正确
	if !security.ComparePassword(userModel.Password, req.OldPassword) {
		log.CtxInfo(ctx, "wrong password")
		return nil, errorx.New(10001, "原密码错误")
	}

	// 检查新旧密码是否相同
	if req.NewPassword == req.OldPassword {
		log.CtxInfo(ctx, "new password cannot be the same as old password")
		return nil, errorx.New(10002, "新密码不能与原密码相同")
	}

	// 检查确认密码是否匹配
	if req.NewPassword != req.ConfirmPassword {
		log.CtxInfo(ctx, "confirm password does not match")
		return nil, errorx.New(10003, "确认密码与新密码不匹配")
	}

	// 生成哈希密码
	if hashPassword, err = security.HashPassword(req.NewPassword); err != nil {
		log.CtxError(ctx, "failed to hash password: %v", err)
		return nil, err
	}

	// 更新密码
	if err = s.UserRepository.UpdatePassword(ctx, userId, hashPassword); err != nil {
		log.CtxError(ctx, "failed to update password: %v", err)
		return nil, err
	}

	return &user.ChangePasswordResp{
		Resp: dto.Success(),
	}, nil
}

func (s *UserService) DeleteAccount(ctx context.Context, req *user.DeleteAccountReq) (*user.DeleteAccountResp, error) {
	var err error
	var userModel *model.User

	// 获取用户ID并转换类型
	userId, ok := ctx.Value(consts.ContextUserID).(primitive.ObjectID)
	if !ok {
		return nil, errors.New("invalid user id in context")
	}

	// 获取用户
	userModel, err = s.UserRepository.FindUserByUserID(ctx, userId)
	if err != nil {
		log.CtxError(ctx, "failed to find user: %v", err)
		return nil, err
	}

	// 校验旧密码是否正确
	if !security.ComparePassword(userModel.Password, req.Password) {
		log.CtxInfo(ctx, "wrong password")
		return nil, errors.New("wrong password")
	}

	if req.Confirmation != "我确认删除账号 "+userModel.Username {
		log.CtxInfo(ctx, "confirmation does not match")
		return nil, errors.New("confirmation does not match")
	}

	if err = s.UserRepository.DeleteUser(ctx, userId); err != nil {
		log.CtxError(ctx, "failed to delete user: %v", err)
		return nil, err
	}

	return &user.DeleteAccountResp{
		Resp: dto.Success(),
	}, nil
}

func (s *UserService) UpdateMyProfile(ctx context.Context, req *user.UpdateMyProfileReq) (*user.UpdateMyProfileResp, error) {
	userId, ok := ctx.Value(consts.ContextUserID).(primitive.ObjectID)
	if !ok {
		return nil, errors.New("invalid user id in context")
	}

	update := bson.M{}
	cnt := 0

	if req.Username != "" {
		update[consts.Username] = req.Username
		cnt++
	}
	if req.Avatar != "" {
		update[consts.Avatar] = req.Avatar
		cnt++
	}
	if req.FirstName != "" {
		update[consts.FirstName] = req.FirstName
		cnt++
	}
	if req.LastName != "" {
		update[consts.LastName] = req.LastName
		cnt++
	}
	if req.Gender != "" {
		update[consts.Gender] = enum.GetUserGenderCode(req.Gender)
		cnt++
	}
	if req.Address != "" {
		update[consts.Address] = req.Address
		cnt++
	}
	if req.Bio != "" {
		update[consts.Bio] = req.Bio
		cnt++
	}
	if req.Birthday != "" {
		t, err := time.Parse("2006-01-02", req.Birthday)
		if err != nil {
			return nil, errors.New("invalid birthday format, use yyyy-MM-dd")
		}
		update[consts.Birthday] = t
		cnt++
	}

	update[consts.UpdatedAt] = time.Now()

	if err := s.UserRepository.UpdateUser(ctx, userId, update); err != nil {
		return nil, err
	}

	return &user.UpdateMyProfileResp{
		Resp:  dto.Success(),
		Count: cnt,
	}, nil
}
