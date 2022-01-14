package model

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uint64
	Username        string        `gorm:"unique;not null"`
	Password        string        `gorm:"not null"`
	Disabled        bool          `gorm:"not null;default:false"`
	IsAdmin         bool          `gorm:"not null;default:false"`
	Groups          []*Group      `gorm:"many2many:groups;"`
	Permissions     []*Permission `gorm:"many2many:permissions;"`
	Icon            string
	BackgroundColor string `gorm:"not null"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[model.(u *User) AfterFind(tx *gorm.DB) (err error)]")
	if u.Password != "" {
		u.Password = ""
	}
	return
}
