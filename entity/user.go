package entity

import (
	"gorm.io/gorm"
)

type Gender int

type Role string

type User struct {
	gorm.Model  `json:"omitEmpty"`
	Username    string   `gorm:"unique;not null;size:15;" json:"username"`
	Password    string   `gorm:"not null" json:"password"`
	Gender      Gender   `gorm:"not null" json:"gender"`
	Role        string   `gorm:"not null" json:"role"`
	AvatarUrl   string   `gorm:"not null" json:"avatar_url"`
	Description string   `gorm:"not null;size:512" json:"description"`
	Valid       bool     `gorm:"not null" json:"valid"`
	Recipes     []Recipe `json:"recipes;omitEmpty"`
}
