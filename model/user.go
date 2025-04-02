package model

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"not null;primaryKey" json:"id"`
	GoogleID   string    `json:"google_id"`
	Name       string    `gorm:"not null" json:"name"`
	Email      string    `gorm:"not null;unique" json:"email"`
	Password   string    `gorm:"not null" json:"password"`
	Role       string    `gorm:"default:User" json:"role"`
	ProfilePic string    `gorm:"default:https://github.com/shadcn.png;" json:"profpic"`
	Blog       []Blog    `gorm:"foreignKey:UserID"`
}
