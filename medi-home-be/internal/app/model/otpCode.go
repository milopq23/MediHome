package model

import "time"

type OTPCode struct {
	OTPID     int64     `gorm:"column:otp_id	"`
	UserID    int64     `gorm:"column:user_id"`
	Code      string    `gorm:"column:code"`
	Purpose   string    `gorm:"column:purpose"`
	ExpiredAt time.Time `gorm:"column:expiredAt"`
	Used      bool      `gorm:"column:used"`
}
