package model

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID     uuid.UUID `gorm:"not null;primaryKey"`
	Name   string    `gorm:"not null;unique" json:"name"`
	ApiKey uuid.UUID `gorm:"not null;unique" json:"api_key"`
	UserID uuid.UUID `gorm:"not null;index;foreignKey:ID,references:users(ID)" json:"user_id"`
	Blog   []Blog    `gorm:"foreignKey:ServiceID"`
}

type Blog struct {
	ID            uuid.UUID `gorm:"not null;primaryKey"`
	Title         string    `gorm:"not null" json:"title"`
	Body          string    `gorm:"not null" json:"body"`
	ServiceID     uuid.UUID `gorm:"not null;index;foreignKey:ID,references:services(ID)" json:"service_id"`
	PublishedDate time.Time
}
