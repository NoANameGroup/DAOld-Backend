package repository

import (
	"context"

	"github.com/NoANameGroup/DAOld-Backend/internal/config"
	"github.com/NoANameGroup/DAOld-Backend/internal/model"
	"github.com/NoANameGroup/DAOld-Backend/pkg/consts"
	"github.com/NoANameGroup/DAOld-Backend/pkg/log"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var _ IElderRepository = (*ElderRepository)(nil)

type IElderRepository interface {
	Insert(ctx context.Context, elder *model.Elder) error

	isFieldExisted(ctx context.Context, field string, value interface{}) (bool, error)
	IsElderExistedByUserID(ctx context.Context, userId bson.ObjectID) (bool, error)

	FindElderByUserID(ctx context.Context, userId bson.ObjectID) (*model.Elder, error)

	DeleteElderByUserID(ctx context.Context, userId bson.ObjectID) error
}

type ElderRepository struct {
	conn *monc.Model
}

func NewElderRepository(config *config.Config) *ElderRepository {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, consts.ElderCollectionName, config.Cache)
	return &ElderRepository{
		conn: conn,
	}
}

func (r *ElderRepository) Insert(ctx context.Context, elder *model.Elder) error {
	if _, err := r.conn.InsertOneNoCache(ctx, elder); err != nil {
		log.CtxError(ctx, "failed to insert elder: %v", err)
		return err
	}

	return nil
}

func (r *ElderRepository) isFieldExisted(ctx context.Context, field string, value interface{}) (bool, error) {
	var err error
	var count int64
	if count, err = r.conn.CountDocuments(ctx, bson.M{field: value}); err != nil {
		log.CtxError(ctx, "failed to check existing %s: %v", field, err)
		return false, err
	}

	return count > 0, nil
}

func (r *ElderRepository) IsElderExistedByUserID(ctx context.Context, userId bson.ObjectID) (bool, error) {
	return r.isFieldExisted(ctx, consts.UserID, userId)
}

func (r *ElderRepository) FindElderByUserID(ctx context.Context, userId bson.ObjectID) (*model.Elder, error) {
	elder := model.Elder{}

	if err := r.conn.FindOneNoCache(ctx, &elder, bson.M{consts.UserID: userId}); err != nil {
		log.CtxError(ctx, "failed to find elder by userId: %v", err)
		return nil, err
	}
	return &elder, nil
}

func (r *ElderRepository) DeleteElderByUserID(ctx context.Context, userId bson.ObjectID) error {
	if _, err := r.conn.DeleteOneNoCache(ctx, bson.M{consts.UserID: userId}); err != nil {
		log.CtxError(ctx, "failed to delete elder by userId: %v", err)
		return err
	}
	return nil
}
