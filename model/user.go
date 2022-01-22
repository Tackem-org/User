package model

import (
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"gorm.io/gorm"
)

type User struct {
	ID              uint64 `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Username        string         `gorm:"unique;not null"`
	Password        string         `gorm:"not null"`
	Disabled        bool           `gorm:"not null;default:false"`
	IsAdmin         bool           `gorm:"not null;default:false"`
	Groups          []Group        `gorm:"many2many:user_groups;"`
	Permissions     []Permission   `gorm:"many2many:user_permissions;"`
	Icon            string         `gorm:"default:'';"`
	BackgroundColor string         `gorm:"not null"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[model.(u *User) AfterFind(tx *gorm.DB) (err error)]")
	if u.Password != "" {
		u.Password = ""
	}
	return
}
