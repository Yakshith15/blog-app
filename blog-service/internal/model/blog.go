package model

import "time"

type Blog struct {
	ID        string    `json:"id"`
	AuthorID  string    `json:"authorId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
