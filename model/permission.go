package model

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID        uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"unique;not null"`
	Users     []*User        `gorm:"many2many:user_permissions;"`
	Groups    []*Group       `gorm:"many2many:group_permissions;"`
}

func AddPermissions(names ...string) {
	for _, name := range names {
		AddPermission(name)
	}
}

func AddPermission(name string) {
	p := Permission{Name: name}
	DB.Where(&p).Find(&p)
	if p.ID == 0 {
		DB.Create(&p)
	}
}
