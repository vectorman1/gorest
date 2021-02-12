package entity

import (
	"database/sql/driver"
	"errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Recipe struct {
	gorm.Model       `json:"omitEmpty"`
	UserID           uint          `json:"user_id"`
	Title            string        `gorm:"not null; size:80" json:"title"`
	ShortDescription string        `gorm:"not null; size:256" json:"short_description"`
	TimeToCookNs     time.Duration `gorm:"not null" json:"time_to_cook_ns"`
	Products         Products      `gorm:"not null;type:text" json:"products"`
	ImageUrl         string        `gorm:"not null" json:"image_url"`
	Description      string        `gorm:"not null;size:2048" json:"description"`
	Tags             Tags          `gorm:"type:text" json:"tags"`
}

type Products []string
type Tags []string

func (p *Products) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("source is not string")
	}
	*p = strings.Split(str, ",")
	return nil
}

func (p Products) Value() (driver.Value, error) {
	if p == nil || len(p) == 0 {
		return nil, nil
	}
	return strings.Join(p, ","), nil
}

func (t *Tags) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("source is not string")
	}
	*t = strings.Split(str, ",")
	return nil
}

func (t Tags) Value() (driver.Value, error) {
	if t == nil || len(t) == 0 {
		return nil, nil
	}
	return strings.Join(t, ","), nil
}
