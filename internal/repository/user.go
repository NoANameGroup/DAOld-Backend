package repository

import (
	"context"
	"errors"
	"github.com/NoANameGroup/DAOld-Backend/consts"
	"github.com/NoANameGroup/DAOld-Backend/internal/config"
	"github.com/NoANameGroup/DAOld-Backend/internal/model"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

const (
	CollectionName = "user"
)

type IUserRepository interface {
	Insert(ctx context.Context, user *model.User) error
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateLastLoginAt(ctx context.Context, userID primitive.ObjectID, t time.Time) error
}

type UserRepository struct {
	conn *monc.Model
}

func NewRepository(config *config.Config) *UserRepository {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.Cache)
	return &UserRepository{
		conn: conn,
	}
}

func (r *UserRepository) Insert(ctx context.Context, user *model.User) error {
	var err error
	// 检查邮箱是否已存在
	var count int64
	if count, err = r.conn.CountDocuments(ctx, bson.M{consts.Email: user.Email}); err != nil {
		log.CtxError(ctx, "failed to check existing email: %v", err)
		return err
	}
	if count > 0 {
		log.CtxError(ctx, "user with email %s already exists", user.Email)
		return errors.New("user with this email already exists")
	}

	// 插入数据库
	if _, err = r.conn.InsertOneNoCache(ctx, user); err != nil {
		log.CtxError(ctx, "failed to insert user: %v", err)
		return err
	}

	return nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var err error
	var user *model.User
	if err = r.conn.FindOneNoCache(ctx, bson.M{consts.Email: email}, &user); err != nil {
		log.CtxError(ctx, "failed to find user by email: %v", err)
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateLastLoginAt(ctx context.Context, userID primitive.ObjectID, t time.Time) error {
	update := bson.M{"$set": bson.M{"lastLoginAt": t}}

	if _, err := r.conn.UpdateByIDNoCache(ctx, userID, update); err != nil {
		log.CtxError(ctx, "failed to update LastLoginAt for user %s: %v", userID.Hex(), err)
		return err
	}

	return nil
}
