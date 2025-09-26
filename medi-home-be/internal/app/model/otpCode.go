package model

import "time"

type OTPCode struct {
	OTPID     int64     `gorm:"column:otp_id,primarykey"`
	UserID    int64     `gorm:"column:user_id"`
	Code      string    `gorm:"column:code"`
	Purpose   string    `gorm:"column:purpose"`
	ExpiredAt time.Time `gorm:"column:expiredAt"`
	Used      bool      `gorm:"column:used"`
}

func (OTPCode) TableName() string {
	return "otp_codes"
}