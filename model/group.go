package model

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	ID          uint64 `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Users       []*User      `gorm:"many2many:user_groups;"`
	Permissions []Permission `gorm:"many2many:group_permissions;"`
}
