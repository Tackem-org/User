package model

import (
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"gorm.io/gorm"
)

type UserRequest struct {
	ID            uint64 `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	RequestUserID uint64         `gorm:"not null"`
	RequestUser   User
	Name          string `gorm:"not null"`
}

func (u *UserRequest) Accept(tx *gorm.DB) (err error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[model.(u *UserRequest) Accept(tx *gorm.DB) (err error)]")
	var user User
	DB.Find(&user, u.RequestUserID)
	DB.Model(&user).Update("Username", u.Name)
	DB.Delete(&u)
	return
}
func (u *UserRequest) Reject(tx *gorm.DB) (err error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[model.(u *UserRequest) Reject(tx *gorm.DB) (err error)]")
	DB.Delete(&u)
	return
}
