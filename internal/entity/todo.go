package entity

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
