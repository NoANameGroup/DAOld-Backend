package service

import (
	"context"
	"github.com/NoANameGroup/DAOld-Backend/internal/dto"
	"github.com/NoANameGroup/DAOld-Backend/internal/dto/session"
	"github.com/NoANameGroup/DAOld-Backend/internal/errorx"
	"github.com/NoANameGroup/DAOld-Backend/internal/jwt"
	"github.com/NoANameGroup/DAOld-Backend/internal/model"
	"github.com/NoANameGroup/DAOld-Backend/internal/repository"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/NoANameGroup/DAOld-Backend/pkg/security"
	"github.com/google/wire"
	"time"
)

type ISessionService interface {
	CreateSession(ctx context.Context, req *session.CreateSessionReq) (*session.CreateSessionResp, error)
	DeleteSession() (*session.DeleteSessionResp, error)
}

type SessionService struct {
	UserRepository *repository.UserRepository
}

var SessionServiceSet = wire.NewSet(
	wire.Struct(new(SessionService), "*"),
	wire.Bind(new(ISessionService), new(*SessionService)),
)

func (s *UserService) CreateSession(ctx context.Context, req *session.CreateSessionReq) (*session.CreateSessionResp, error) {
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
		return nil, errorx.ErrUsernameOrPasswordIncorrect
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

	return &session.CreateSessionResp{
		Resp:        dto.Success(),
		AccessToken: token,
		UserID:      newUser.ID,
	}, nil
}

func (s *UserService) DeleteSession() (*session.DeleteSessionResp, error) {
	return &session.DeleteSessionResp{
		Resp: dto.Success(),
	}, nil
}
