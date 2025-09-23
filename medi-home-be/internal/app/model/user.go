package model

import (
	"time"
)

type User struct{
	UserID int64 ``
	Email string
	Password string
	Phone string
	Name string
	Gender string
	Avatar string
	Point int64
	IsVerified bool
	CreatedAt time.Time
	UpdatedAt time.Time
}