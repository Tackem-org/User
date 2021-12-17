package model

import (
	"os"
	"time"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/User/flags"
	"github.com/Tackem-org/User/password"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func Setup() {

	password.SetupSalt()
	setupDB()

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Group{})
	DB.AutoMigrate(&Permission{})

	var count int64
	DB.Model(&User{}).Count(&count)
	if count == 0 {
		f, _ := os.Create("/config/adminpassword")
		newPassword := helpers.RandStr(8)
		f.WriteString(newPassword)
		f.Close()

		DB.Create(&User{
			Model:       gorm.Model{},
			Username:    "admin",
			Password:    password.Hash(newPassword),
			Disabled:    false,
			IsAdmin:     true,
			Groups:      []*Group{},
			Permissions: []*Permission{},
		})

	}
}

func setupDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(*flags.DatabaseFile), &gorm.Config{
		Logger: logger.New(
			logging.CustomLogger("GORM"),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	})
	if err != nil {
		panic("failed to Open database")
	}
}
