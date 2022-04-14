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
