package domain

import "time"

type OTPCode struct {
	ID        int
	Phone     string
	Code      string
	ExpiresAt time.Time
	Used      bool
	CreatedAt time.Time
}
