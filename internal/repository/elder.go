package repository

import (
	"context"

	"github.com/NoANameGroup/DAOld-Backend/internal/config"
	"github.com/NoANameGroup/DAOld-Backend/internal/model"
	"github.com/NoANameGroup/DAOld-Backend/pkg/consts"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ IElderRepository = (*ElderRepository)(nil)

type IElderRepository interface {
	Insert(ctx context.Context, elder *model.Elder) error
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
		return err
	}

	return nil
}
