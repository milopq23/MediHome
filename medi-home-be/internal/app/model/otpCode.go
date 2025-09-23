package model

import "time"

type OTPCode struct{
	OTPID int64
	UserID int64 
	Code string
	Purpose string
	ExpiredAt time.Time
	Used bool
}