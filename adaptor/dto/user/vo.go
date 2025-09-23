package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserVO struct {
	ID     primitive.ObjectID `json:"id"`
	Name   string             `json:"name"`
	Email  string             `json:"email"`
	Phone  string             `json:"phone"`
	Avatar string             `json:"avatar"`
}
