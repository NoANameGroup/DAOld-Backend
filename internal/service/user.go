package service

import (
	"context"
	"errors"
	"github.com/NoANameGroup/DAOld-Backend/internal/dto"
	"github.com/NoANameGroup/DAOld-Backend/internal/dto/user"
	"github.com/NoANameGroup/DAOld-Backend/internal/jwt"
	"github.com/NoANameGroup/DAOld-Backend/internal/model"
	"github.com/NoANameGroup/DAOld-Backend/internal/repository"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/NoANameGroup/DAOld-Backend/pkg/security"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IUserService interface {
	Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterResp, error)
	Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error)
}

type UserService struct {
	UserMapper *repository.UserRepository
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
		Username:  req.Name,
		Password:  hashPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 插入数据库
	if err = s.UserMapper.Insert(ctx, newUser); err != nil {
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
	if newUser, err = s.UserMapper.FindUserByEmail(ctx, req.Email); err != nil {
		log.CtxError(ctx, "failed to find user: %v", err)
		return nil, err
	}

	// 校验密码是否正确
	if !security.ComparePassword(req.Password, newUser.Password) {
		log.CtxInfo(ctx, "username or password incorrect")
		return nil, errors.New("username or password incorrect")
	}

	// 更新最后登录时间
	if err = s.UserMapper.UpdateLastLoginAt(ctx, newUser.ID, time.Now()); err != nil {
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
