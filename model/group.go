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
	Name        string         `gorm:"unique;not null"`
	Users       []*User        `gorm:"many2many:user_groups;"`
	Permissions []Permission   `gorm:"many2many:group_permissions;"`
}

func AddGroups(names ...string) {
	for _, name := range names {
		AddGroup(name)
	}
}

func AddGroup(name string) {
	g := Group{Name: name}
	DB.Where(&g).Find(&g)
	if g.ID == 0 {
		DB.Create(&g)
	}
}
