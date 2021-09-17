package entity

import (
	"time"
)

// Item struct
type Item struct {
	ID          int       `json:"id,primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url"`
	AutoUpdate  bool      `json:"auto_update,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
