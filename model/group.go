package model

import (
	"errors"
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

func AddGroup(name string) {
	g := Group{Name: name}
	if err := DB.Where(&g).Find(&g).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			DB.Create(&g)
		}
	} else if g.ID == 0 {
		DB.Create(&g)
	}
}
