package repository

import (
	"context"
	"time"

	"github.com/NoANameGroup/DAOld-Backend/internal/config"
	"github.com/NoANameGroup/DAOld-Backend/internal/model"
	"github.com/NoANameGroup/DAOld-Backend/pkg/consts"
	"github.com/NoANameGroup/DAOld-Backend/pkg/consts/enum"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var _ IUserRepository = (*UserRepository)(nil)

type IUserRepository interface {
	isFieldExisted(ctx context.Context, field string, value interface{}) (bool, error)
	IsEmailExisted(ctx context.Context, email string) (bool, error)
	IsPhoneExisted(ctx context.Context, phone string) (bool, error)
	IsUsernameExisted(ctx context.Context, username string) (bool, error)

	Insert(ctx context.Context, user *model.User) error
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	FindUserByUserID(ctx context.Context, userId bson.ObjectID) (*model.User, error)
	DeleteUserByUserID(ctx context.Context, userId bson.ObjectID) error
	UpdateUserByUserID(ctx context.Context, userId bson.ObjectID, update bson.M) error
	IsAdmin(ctx context.Context, userId bson.ObjectID) (bool, error)

	updateFieldByUserID(ctx context.Context, userId bson.ObjectID, update bson.M) error
	UpdateLastLoginAtByUserID(ctx context.Context, userId bson.ObjectID, t time.Time) error
	UpdatePasswordByUserID(ctx context.Context, userId bson.ObjectID, password string) error
	UpdateUserRoleByUserID(ctx context.Context, userId bson.ObjectID, role enum.UserRole) error
	UpdateEmailByUserID(ctx context.Context, userId bson.ObjectID, email string) error
	UpdatePhoneByUserID(ctx context.Context, userId bson.ObjectID, phone string) error
	UpdateUsernameByUserID(ctx context.Context, userId bson.ObjectID, username string) error
	UpdateUpdatedAtByUserID(ctx context.Context, userId bson.ObjectID, t time.Time) error
}

type UserRepository struct {
	conn *monc.Model
}

func NewUserRepository(config *config.Config) *UserRepository {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, consts.UserCollectionName, config.Cache)
	return &UserRepository{
		conn: conn,
	}
}

func (r *UserRepository) isFieldExisted(ctx context.Context, field string, value interface{}) (bool, error) {
	var err error
	var count int64
	if count, err = r.conn.CountDocuments(ctx, bson.M{field: value}); err != nil {
		log.CtxError(ctx, "failed to check existing %s: %v", field, err)
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepository) IsEmailExisted(ctx context.Context, email string) (bool, error) {
	return r.isFieldExisted(ctx, consts.Email, email)
}

func (r *UserRepository) IsPhoneExisted(ctx context.Context, phone string) (bool, error) {
	return r.isFieldExisted(ctx, consts.Phone, phone)
}

func (r *UserRepository) IsUsernameExisted(ctx context.Context, username string) (bool, error) {
	return r.isFieldExisted(ctx, consts.Username, username)
}

func (r *UserRepository) Insert(ctx context.Context, user *model.User) error {
	if _, err := r.conn.InsertOneNoCache(ctx, user); err != nil {
		log.CtxError(ctx, "failed to insert user: %v", err)
		return err
	}

	return nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}

	if err := r.conn.FindOneNoCache(ctx, &user, bson.M{consts.Email: email}); err != nil {
		log.CtxError(ctx, "failed to find user by email: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindUserByUserID(ctx context.Context, userId bson.ObjectID) (*model.User, error) {
	user := model.User{}

	if err := r.conn.FindOneNoCache(ctx, &user, bson.M{consts.ID: userId}); err != nil {
		log.CtxError(ctx, "failed to find user by userId: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) DeleteUserByUserID(ctx context.Context, userId bson.ObjectID) error {
	if _, err := r.conn.DeleteOneNoCache(ctx, bson.M{consts.ID: userId}); err != nil {
		log.CtxError(ctx, "failed to delete user %s: %v", userId.Hex(), err)
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUserByUserID(ctx context.Context, userId bson.ObjectID, update bson.M) error {
	if _, err := r.conn.UpdateByIDNoCache(ctx, userId, bson.M{"$set": update}); err != nil {
		log.CtxError(ctx, "failed to update user %s: %v", userId.Hex(), err)
		return err
	}

	return nil
}

func (r *UserRepository) IsAdmin(ctx context.Context, userId bson.ObjectID) (bool, error) {
	user, err := r.FindUserByUserID(ctx, userId)
	if err != nil {
		log.CtxError(ctx, "failed to find user by userId: %v", err)
		return false, err
	}

	return user.Role == enum.RoleAdmin, nil
}

func (r *UserRepository) updateFieldByUserID(ctx context.Context, userId bson.ObjectID, update bson.M) error {
	if _, err := r.conn.UpdateByIDNoCache(ctx, userId, bson.M{"$set": update}); err != nil {
		log.CtxError(ctx, "failed to update user %s: %v", userId.Hex(), err)
		return err
	}

	return nil
}

func (r *UserRepository) UpdateLastLoginAtByUserID(ctx context.Context, userId bson.ObjectID, t time.Time) error {
	return r.updateFieldByUserID(ctx, userId, bson.M{consts.LastLoginAt: t})
}

func (r *UserRepository) UpdatePasswordByUserID(ctx context.Context, userId bson.ObjectID, hashPassword string) error {
	return r.updateFieldByUserID(ctx, userId, bson.M{consts.Password: hashPassword})
}

func (r *UserRepository) UpdateUserRoleByUserID(ctx context.Context, userId bson.ObjectID, role enum.UserRole) error {
	return r.updateFieldByUserID(ctx, userId, bson.M{consts.Role: role})
}

func (r *UserRepository) UpdateEmailByUserID(ctx context.Context, userId bson.ObjectID, email string) error {
	return r.updateFieldByUserID(ctx, userId, bson.M{consts.Email: email})
}

func (r *UserRepository) UpdatePhoneByUserID(ctx context.Context, userId bson.ObjectID, phone string) error {
	return r.updateFieldByUserID(ctx, userId, bson.M{consts.Phone: phone})
}

func (r *UserRepository) UpdateUsernameByUserID(ctx context.Context, userId bson.ObjectID, username string) error {
	return r.updateFieldByUserID(ctx, userId, bson.M{consts.Username: username})
}

func (r *UserRepository) UpdateUpdatedAtByUserID(ctx context.Context, userId bson.ObjectID, t time.Time) error {
	return r.updateFieldByUserID(ctx, userId, bson.M{consts.UpdatedAt: t})
}
