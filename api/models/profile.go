package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID              string         `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	UserID          string         `gorm:"type:uuid;not null;unique" json:"user_id"`
	LastVisit       time.Time      `gorm:"type:timestamp" json:"last_visit"`
	InvitationCount int            `gorm:"default:0" json:"invitation_count"`
	IRCKey          string         `gorm:"type:varchar(255)" json:"irc_key"`
	RSSKey          string         `gorm:"type:varchar(255)" json:"rss_key"`
	User            User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt       time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
