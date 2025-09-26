package util

import (
	"fmt"
	"os"
	"gopkg.in/gomail.v2"
)

func SendEmailOTP(toEmail, otp string) error {
	from := os.Getenv("OTP_EMAIL")      
	password := os.Getenv("OTP_PASSWORD")      
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s\nIt will expire in 5 minutes.", otp))

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	return d.DialAndSend(m)
}