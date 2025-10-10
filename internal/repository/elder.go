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
	FindElderByUserID(ctx context.Context, userId bson.ObjectID) (*model.Elder, error)
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

func (r *ElderRepository) FindElderByUserID(ctx context.Context, userId bson.ObjectID) (*model.Elder, error) {
	elder := model.Elder{}

	if err := r.conn.FindOneNoCache(ctx, &elder, bson.M{consts.UserID: userId}); err != nil {
		log.CtxError(ctx, "failed to find elder by userId: %v", err)
		return nil, err
	}
	return &elder, nil
}
