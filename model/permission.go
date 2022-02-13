package model

import (
	"errors"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
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

func AddPermission(name string) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[model.AddPermission(name string)] {name=%s}", name)
	p := Permission{Name: name}
	if err := DB.Where(&p).Find(&p).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			DB.Create(&p)
		}
	} else if p.ID == 0 {
		DB.Create(&p)
	}
}
