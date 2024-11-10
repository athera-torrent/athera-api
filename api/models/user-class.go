package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	RootUserClass      = "root"
	AdminUserClass     = "admin"
	ModeratorUserClass = "moderator"
	DefaultUserClass   = "user"
)

type UserClass struct {
	ID        string         `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    string         `gorm:"type:uuid;not null;unique" json:"user_id"`
	Type      string         `gorm:"type:varchar(50);not null" json:"type"`
	StartDate time.Time      `gorm:"type:timestamp;not null" json:"start_date"`
	EndDate   *time.Time     `gorm:"type:timestamp" json:"end_date,omitempty"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	Features  []string       `gorm:"type:text[]" json:"features"`
	User      User           `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
