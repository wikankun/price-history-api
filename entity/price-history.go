package entity

import "time"

// PriceHistory struct
type PriceHistory struct {
	ID        int       `json:"id,primary_key"`
	ItemID    int       `json:"item_id"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Item      Item      `json:",omitempty"`
}
