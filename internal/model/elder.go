package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Elder struct {
	ID                primitive.ObjectID `bson:"id"`
	UserID            primitive.ObjectID `bson:"userId"`
	BlockChainAddress string             `bson:"blockChainAddress"`
	Balance           float64            `bson:"balance"`
}
