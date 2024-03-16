package domain

import "time"

type OtpInfo struct {
	ID         uint `gorm:"primaryKey"`
	Email      string
	OTP        int
	Expiration time.Time
}
