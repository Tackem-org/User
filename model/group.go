package model

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name        string
	Users       []*User       `gorm:"many2many:users;"`
	SubGroups   []*Group      `gorm:"many2many:sub_groups;"`
	Permissions []*Permission `gorm:"many2many:permissions;"`
}
