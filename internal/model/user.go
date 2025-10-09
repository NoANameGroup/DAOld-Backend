package model

import (
	"time"

	"github.com/NoANameGroup/DAOld-Backend/pkg/consts/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username"`
	FirstName   string             `bson:"firstName"`
	LastName    string             `bson:"lastName"`
	IDCard      string             `bson:"idCard"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	Phone       string             `bson:"phone"`
	Avatar      string             `bson:"avatar"`
	Address     string             `bson:"address"`
	Role        enum.UserRole      `bson:"role"`
	Status      enum.UserStatus    `bson:"status"`
	Gender      enum.UserGender    `bson:"gender"`
	Birthday    time.Time          `bson:"birthday"`
	Bio         string             `bson:"bio"`
	LastLoginAt time.Time          `bson:"lastLoginAt"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}
