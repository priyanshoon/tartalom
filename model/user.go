package model

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	GoogleID   string    `json:"google_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	ProfilePic string    `gorm:"default:https://github.com/shadcn.png;" json:"profpic"`
	Blog       []Blog
}
