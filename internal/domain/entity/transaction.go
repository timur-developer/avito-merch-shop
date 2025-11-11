package entity

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	FromUserID      int       `json:"from_user_id,omitempty"`
	ToUserID        int       `json:"to_user_id,omitempty"`
	Amount          int       `json:"amount"`
	TransactionType string    `json:"transaction_type"` // transfer или purchase
	CreatedAt       time.Time `json:"created_at"`
}
