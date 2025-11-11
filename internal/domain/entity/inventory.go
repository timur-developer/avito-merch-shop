package entity

import "time"

type InventoryItem struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	ItemID      int       `json:"item_id"`
	Quantity    int       `json:"quantity"`
	PurchasedAt time.Time `json:"purchased_at"`
}
