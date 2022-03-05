package model

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID          uint64 `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Username    string         `gorm:"unique;not null"`
	Password    string         `gorm:"not null"`
	Disabled    bool           `gorm:"not null;default:false"`
	IsAdmin     bool           `gorm:"not null;default:false"`
	Groups      []Group        `gorm:"many2many:user_groups;"`
	Permissions []Permission   `gorm:"many2many:user_permissions;"`
	Icon        string         `gorm:"default:'';"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.Password != "" {
		u.Password = ""
	}
	return
}

func (u *User) AllPermissionStrings() []string {
	var s []string
	for _, p := range u.Permissions {
		s = append(s, p.Name)
	}
	var group Group
	for _, g := range u.Groups {
		DB.Preload(clause.Associations).First(&group, g.ID)
		for _, p := range group.Permissions {
			s = append(s, p.Name)

		}
	}
	return s
}
