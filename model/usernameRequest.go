package model

import (
	"time"

	"gorm.io/gorm"
)

type UsernameRequest struct {
	ID            uint64 `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	RequestUserID uint64         `gorm:"not null"`
	RequestUser   User
	Name          string `gorm:"not null"`
}

func (u *UsernameRequest) Accept(tx *gorm.DB) (err error) {
	var user User
	DB.Find(&user, u.RequestUserID)
	DB.Model(&user).Update("Username", u.Name)
	DB.Delete(&u)
	return
}
func (u *UsernameRequest) Reject(tx *gorm.DB) (err error) {
	DB.Delete(&u)
	return
}
