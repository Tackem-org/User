package model

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Name   string
	Users  []*User  `gorm:"many2many:users_permissions;"`
	Groups []*Group `gorm:"many2many:groups_permissions;"`
}
