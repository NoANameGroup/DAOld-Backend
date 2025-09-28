package user

import (
	"time"
)

type UserVO struct {
	Username    string    `json:"username"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Avatar      string    `json:"avatar"`
	Address     string    `json:"address"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	Gender      string    `json:"gender"`
	Birthday    string    `json:"birthday"`
	Bio         string    `json:"bio"`
	LastLoginAt time.Time `json:"lastLoginAt"`
	CreatedAt   time.Time `json:"createdAt"`
}
