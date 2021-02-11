package models

import "gorest/entity"

type UserRequest struct {
	Username    string
	Password    string
	Gender      entity.Gender
	Role        string
	AvatarUrl   string
	Description string
	Valid       bool
	Recipes     []entity.Recipe `json:"omitEmpty"`
}
