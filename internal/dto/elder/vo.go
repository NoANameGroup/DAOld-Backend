package elder

import "go.mongodb.org/mongo-driver/bson/primitive"

type ElderVO struct {
	ID                primitive.ObjectID `json:"id"`
	UserID            primitive.ObjectID `json:"userId"`
	BlockChainAddress string             `json:"blockChainAddress"`
	Balance           float64            `json:"balance"`
}
