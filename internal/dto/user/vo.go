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
	Role        int       `json:"role"`
	Status      int       `json:"status"`
	Gender      int       `json:"gender"`
	Birthday    time.Time `json:"birthday"`
	Bio         string    `json:"bio"`
	LastLoginAt time.Time `json:"lastLoginAt"`
	CreatedAt   time.Time `json:"createdAt"`
}
