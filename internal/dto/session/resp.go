package session

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateSessionResp struct {
	*dto.Resp
	UserID      primitive.ObjectID `json:"userId"`
	AccessToken string             `json:"accessToken"`
}

type DeleteSessionResp struct {
	*dto.Resp
}
