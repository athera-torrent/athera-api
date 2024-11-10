package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                          string          `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Username                    string          `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	Email                       string          `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash                string          `gorm:"type:varchar(255);not null" json:"-"`
	IsVerified                  bool            `gorm:"default:false" json:"is_verified"`
	VerificationToken           string          `gorm:"type:varchar(255)" json:"-"`
	VerificationTokenExpiredAt  time.Time       `gorm:"type:timestamp" json:"-"`
	ResetPasswordToken          string          `gorm:"type:varchar(255)" json:"-"`
	ResetPasswordTokenExpiredAt time.Time       `gorm:"type:timestamp" json:"-"`
	CreatedAt                   time.Time       `gorm:"<-:create" json:"created_at"`
	UpdatedAt                   time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt                   gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
	IsInvited                   bool            `gorm:"default:false" json:"is_invited"`
	Invites                     []Invite        `gorm:"foreignKey:InvitedBy" json:"invites,omitempty"`
	Profile                     *Profile        `gorm:"foreignKey:UserID" json:"profile,omitempty"`
	UserClass                   *UserClass      `gorm:"foreignKey:UserID" json:"user_class,omitempty"`
	APIKeys                     []APIAccessKey  `gorm:"foreignKey:UserID" json:"api_keys,omitempty"`
	TorrentPassKey              *TorrentPassKey `gorm:"foreignKey:UserID" json:"torrent_pass_key,omitempty"`
}
