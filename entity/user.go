package entity

import (
	"database/sql"
	"time"
)

type Gender int
type Role string

type User struct {
	ID          uint         `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time    `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `gorm:"index" json:"deleted_at"`
	Username    string       `gorm:"unique;not null;size:15;" json:"username"`
	Password    string       `gorm:"not null" json:"password"`
	Gender      Gender       `gorm:"not null" json:"gender"`
	Role        string       `gorm:"not null;default:user" json:"role"`
	AvatarUrl   string       `gorm:"not null" json:"avatar_url"`
	Description string       `gorm:"not null;size:512" json:"description"`
	Valid       bool         `gorm:"not null" json:"valid"`
	Recipes     Recipes      `gorm:"ForeignKey:UserID;OnDelete:CASCADE" json:"-"`
}

type Recipes []Recipe
