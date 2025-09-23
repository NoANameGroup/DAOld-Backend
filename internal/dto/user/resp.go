package user

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterResp struct {
	*dto.Resp
}

type LoginResp struct {
	*dto.Resp
	UserID      primitive.ObjectID `json:"userId"`
	AccessToken string             `json:"accessToken"`
}
