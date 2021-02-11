package entity

import (
	"gorm.io/gorm"
)

type Gender int

const (
	GenderMale Gender = iota
	GenderFemale
	GenderOther
)

type Role string

const (
	Admin  Role = "admin"
	Normal Role = "user"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique;not null;size:15;"`
	Password    string `gorm:"not null"`
	Gender      Gender `gorm:"not null"`
	Role        string `gorm:"not null"`
	AvatarUrl   string `gorm:"not null"`
	Description string `gorm:"not null;size:512"`
	Valid       bool   `gorm:"not null"`
	Recipes     []Recipe
}
