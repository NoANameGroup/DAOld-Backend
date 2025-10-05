package session

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/dto"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateSessionResp struct {
	*dto.Resp
	UserID      bson.ObjectID `json:"userId"`
	AccessToken string        `json:"accessToken"`
}

type DeleteSessionResp struct {
	*dto.Resp
}
