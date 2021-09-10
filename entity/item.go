package entity

import "time"

type Item struct {
	ID          int       `json:"id,primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Url         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
