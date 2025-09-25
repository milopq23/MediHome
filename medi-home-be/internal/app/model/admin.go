package model

import "time"

type Admin struct{
	AdminID int64 `gorm:"admin_id;primarykey" json:""`
	Email string
	Password string
	Phone string
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}