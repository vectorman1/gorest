package entity

import (
	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
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
	Products         pq.StringArray `gorm:"not null;type:text" json:"products"`
	ImageUrl         string         `gorm:"not null" json:"image_url"`
	Description      string         `gorm:"not null;size:2048" json:"description"`
	Tags             pq.StringArray `gorm:"type:text" json:"tags"`

	User   User `json:"-"`
	UserID uint `gorm:"not null" json:"user_id"`
}

type Products pq.StringArray
type Tags pq.StringArray
