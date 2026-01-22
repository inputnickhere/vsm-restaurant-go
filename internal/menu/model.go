package menu

import "time"

type Item struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
