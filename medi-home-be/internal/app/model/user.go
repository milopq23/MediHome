package model

import (
	"time"
)

type User struct {
	UserID     int64     `gorm:"primarykey,column:user_id,autoIncrement" json:"id"`
	Email      string    `gorm:"column:email;unique;not null" json:"email"`
	Password   string    `gorm:"column:password" json:"-"`
	Phone      string    `gorm:"column:phone;unique" json:"phone"`
	Name       string    `gorm:"column:name" json:"name"`
	Gender     string    `gorm:"column:gender" json:"gender"`
	Avatar     string    `gorm:"column:avatar" json:"avatar"`
	Point      int64     `gorm:"column:point;default:0"`
	IsVerified bool      `gorm:"column:is_verified"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (User) TableName() string{
	return "users"
}