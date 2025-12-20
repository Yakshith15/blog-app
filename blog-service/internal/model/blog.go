package model

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	ID        uuid.UUID `json:"id"`
	AuthorID  uuid.UUID `json:"authorId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
