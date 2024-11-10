package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	APIKeyStatusActive  = "active"
	APIKeyStatusRevoked = "revoked"
	APIKeyStatusExpired = "expired"
	PermissionReadOnly  = "read"
	PermissionWrite     = "write"
	PermissionAdmin     = "admin"
)

type APIAccessKey struct {
	ID          string         `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	UserID      string         `gorm:"type:uuid;not null;index" json:"user_id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Key         string         `gorm:"type:varchar(255);unique;not null" json:"-"`
	Status      string         `gorm:"type:varchar(20);default:'active'" json:"status"`
	LastUsedAt  *time.Time     `gorm:"type:timestamp" json:"last_used_at,omitempty"`
	ExpiresAt   *time.Time     `gorm:"type:timestamp" json:"expires_at,omitempty"`
	Permissions []string       `gorm:"type:text[]" json:"permissions"`
	User        User           `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt   time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
