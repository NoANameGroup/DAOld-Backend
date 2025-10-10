package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Elder struct {
	ID                bson.ObjectID `bson:"_id"`
	UserID            bson.ObjectID `bson:"userId"`
	BlockChainAddress string        `bson:"blockChainAddress"`
	Balance           float64       `bson:"balance"`
	CreatedAt         time.Time     `bson:"createdAt"`
	UpdatedAt         time.Time     `bson:"updatedAt"`
}
