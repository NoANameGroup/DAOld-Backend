package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username"`
	FirstName   string             `bson:"firstName"`
	LastName    string             `bson:"lastName"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	Phone       string             `bson:"phone"`
	Avatar      string             `bson:"avatar"`
	Address     string             `bson:"address"`
	Rule        int                `bson:"rule"`
	Status      int                `bson:"status"`
	Gender      int                `bson:"gender"`
	Birthday    time.Time          `bson:"birthday"`
	Bio         string             `bson:"bio"`
	LastLoginAt time.Time          `bson:"lastLoginAt"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}
