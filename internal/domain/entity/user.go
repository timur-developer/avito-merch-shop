package entity

import "time"

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string
	Coins        int       `json:"coins"`
	CreatedAt    time.Time `json:"created_at"`
}
