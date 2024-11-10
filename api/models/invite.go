package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	InviteStatusPending  = "pending"
	InviteStatusAccepted = "accepted"
	InviteStatusExpired  = "expired"
)

type Invite struct {
	ID        string         `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Email     string         `gorm:"type:varchar(255);not null" json:"email"`
	Token     string         `gorm:"type:varchar(255);unique;not null" json:"-"`
	Status    string         `gorm:"type:varchar(20);default:'pending'" json:"status"`
	ExpiresAt time.Time      `gorm:"type:timestamp;not null" json:"expires_at"`
	InvitedBy string         `gorm:"type:uuid;not null" json:"invited_by"`
	User      *User          `gorm:"foreignKey:InvitedBy" json:"user,omitempty"`
	CreatedAt time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
