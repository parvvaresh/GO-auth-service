package domain

import "time"

type User struct {
	ID              int       `json:"id"`
	Phone           string    `json:"phone"`
	Username        string    `json:"username"`
	PasswordHash    string    `json:"-"`
	IsPhoneVerified bool      `json:"is_phone_verified"`
	CreatedAt       time.Time `json:"created_at"`
}
