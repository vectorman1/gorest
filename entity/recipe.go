package entity

import (
	"database/sql/driver"
	"github.com/dystopia-systems/alaskalog"
	"github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type Recipe struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        mysql.NullTime `gorm:"index" json:"deleted_at"`
	Title            string         `gorm:"not null; size:80" json:"title"`
	ShortDescription string         `gorm:"not null; size:256" json:"short_description"`
	TimeToCookNs     time.Duration  `gorm:"not null" json:"time_to_cook_ns"`
	Products         Products       `gorm:"not null;type:text" json:"products"`
	ImageUrl         string         `gorm:"not null" json:"image_url"`
	Description      string         `gorm:"not null;size:2048" json:"description"`
	Tags             Tags           `gorm:"type:text" json:"tags"`
	User             User           `json:"-"`
	UserID           uint           `gorm:"not null" json:"-"`
}

type Products []string
type Tags []string

func (p *Products) Scan(src interface{}) error {
	str, _ := src.(string)
	*p = strings.Split(str, ",")
	alaskalog.Logger.Println(*p)
	return nil
}

func (p Products) Value() (driver.Value, error) {
	if p == nil || len(p) == 0 {
		return nil, nil
	}
	return strings.Join(p, ","), nil
}

func (t *Tags) Scan(src interface{}) error {
	str, _ := src.(string)
	*t = strings.Split(str, ",")
	return nil
}

func (t Tags) Value() (driver.Value, error) {
	if t == nil || len(t) == 0 {
		return nil, nil
	}
	return strings.Join(t, ","), nil
}
