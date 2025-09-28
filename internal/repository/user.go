package repository

import (
	"context"
	"errors"
	"github.com/NoANameGroup/DAOld-Backend/internal/config"
	"github.com/NoANameGroup/DAOld-Backend/internal/consts"
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
	UpdateLastLoginAt(ctx context.Context, userId primitive.ObjectID, t time.Time) error
	FindUserByUserID(ctx context.Context, userId primitive.ObjectID)
	UpdatePassword(ctx context.Context, userId primitive.ObjectID, password string) error
	DeleteUser(ctx context.Context, userId primitive.ObjectID) error
	UpdateUser(ctx context.Context, userId primitive.ObjectID, update bson.M) error
}

type UserRepository struct {
	conn *monc.Model
}

func NewUserRepository(config *config.Config) *UserRepository {
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
	user := model.User{}
	log.CtxInfo(ctx, "FindUserByEmail in collection=%s, filter=%+v", CollectionName, bson.M{consts.Email: email})

	if err = r.conn.FindOneNoCache(ctx, &user, bson.M{consts.Email: email}); err != nil {
		log.CtxError(ctx, "failed to find user by email: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateLastLoginAt(ctx context.Context, userId primitive.ObjectID, t time.Time) error {
	if _, err := r.conn.UpdateByIDNoCache(ctx, userId, bson.M{"$set": bson.M{consts.LastLoginAt: t}}); err != nil {
		log.CtxError(ctx, "failed to update LastLoginAt for user %s: %v", userId.Hex(), err)
		return err
	}

	return nil
}

func (r *UserRepository) FindUserByUserID(ctx context.Context, userId primitive.ObjectID) (*model.User, error) {
	var err error
	user := model.User{}
	log.CtxInfo(ctx, "FindUserByUserID in collection=%s, filter=%+v", CollectionName, bson.M{consts.UserID: userId})

	if err = r.conn.FindOneNoCache(ctx, &user, bson.M{consts.ID: userId}); err != nil {
		log.CtxError(ctx, "failed to find user by userId: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdatePassword(ctx context.Context, userId primitive.ObjectID, hashPassword string) error {
	if _, err := r.conn.UpdateByIDNoCache(ctx, userId, bson.M{"$set": bson.M{consts.Password: hashPassword}}); err != nil {
		log.CtxError(ctx, "failed to update password for user %s: %v", userId.Hex(), err)
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, userId primitive.ObjectID) error {
	if _, err := r.conn.DeleteOneNoCache(ctx, bson.M{consts.ID: userId}); err != nil {
		log.CtxError(ctx, "failed to delete user %s: %v", userId.Hex(), err)
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, userId primitive.ObjectID, update bson.M) error {
	if _, err := r.conn.UpdateByIDNoCache(ctx, userId, bson.M{"$set": update}); err != nil {
		log.CtxError(ctx, "failed to update user %s: %v", userId.Hex(), err)
		return err
	}
	return nil
}
