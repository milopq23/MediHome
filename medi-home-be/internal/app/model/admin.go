package model

import "time"

type Admin struct{
	AdminID int64
	Email string
	Password string
	Phone string
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}