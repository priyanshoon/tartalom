package model

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Service_ID uuid.UUID
	Name       string
}

type Blog struct {
	Blog_ID       uuid.UUID `json:"blog_id"`
	Title         string    `gorm:"not null" json:"title"`
	Body          string    `gorm:"not null" json:"body"`
	UserID        uuid.UUID `gorm:"not null;index;foreignKey:ID,references:users(ID)" json:"user_id"`
	PublishedDate time.Time
}
