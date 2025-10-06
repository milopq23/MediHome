package model

import "time"

type Admin struct {
	AdminID   int64     `gorm:"column:admin_id;primarykey;autoIncrement"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"column:email;unique;not null" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Role      string    `gorm:"column:role" json:"role"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime"`
}

func (Admin) TableName() string {
	return "admins"
}
