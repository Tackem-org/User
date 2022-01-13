package model

import (
	"os"
	"time"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/User/password"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func Setup(dbFile string) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[model.Setup(dbFile string)] {dbFile=%s}", dbFile)
	password.SetupSalt()
	var err error
	DB, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{
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
