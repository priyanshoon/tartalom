package model

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	Blog_ID       string `json:"blog_id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	PublishedDate time.Time
	UserID        uuid.UUID
}
