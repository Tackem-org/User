package model

import (
	"os"
	"time"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/User/password"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func Setup(dbFile string) {
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

	if err := DB.AutoMigrate(&Permission{}, &Group{}, &User{}, &UsernameRequest{}); err != nil {
		logging.Fatal("unable autoMigrateDB - " + err.Error())
	}

	var count int64
	DB.Model(&User{}).Count(&count)
	if count == 0 {
		f, _ := os.Create("/config/adminpassword")
		newPassword := helpers.RandStr(8)
		f.WriteString(newPassword)
		f.Close()

		DB.Create(&User{
			Username: "admin",
			Password: password.Hash(newPassword),
			Disabled: false,
			IsAdmin:  true,
			Icon:     "",
		})
		DB.Create(&User{
			Username: "user",
			Password: password.Hash("user"),
			Disabled: false,
			IsAdmin:  false,
			Icon:     "",
		})

	}

	//Master Permissions
	AddPermission("do_tasks")

	//User Permissions
	AddPermission("system_user_change_own_password")
	AddPermission("system_user_change_own_username")
	AddPermission("system_user_request_change_of_username")
	AddPermission("system_user_action_change_of_username")

	//User Groups
	AddGroup("user")
	AddGroup("super_user")
	AddGroup("power_user")

}
