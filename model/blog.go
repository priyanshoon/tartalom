package model

import "github.com/google/uuid"

type Blog struct {
	Blog_ID string `json:"blog_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	UserID  uuid.UUID
}
