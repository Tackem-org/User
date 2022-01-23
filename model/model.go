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

	if err := DB.AutoMigrate(&Permission{}, &Group{}, &User{}); err != nil {
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
			Username:        "admin",
			Password:        password.Hash(newPassword),
			Disabled:        false,
			IsAdmin:         true,
			Icon:            "",
			BackgroundColor: "#160686",
		})
		DB.Create(&User{
			Username:        "user",
			Password:        password.Hash("user"),
			Disabled:        false,
			IsAdmin:         false,
			Icon:            "",
			BackgroundColor: "#160686",
		})

	}
	DB.Model(&Permission{}).Count(&count)
	if count == 0 {
		p := []Permission{
			{Name: "system_user_view_own_user_profile"},
			{Name: "system_user_view_other_user_profile"},
			{Name: "system_user_edit_own_user_profile"},
			{Name: "system_user_edit_other_user_profile"},

			{Name: "system_user_edit_group_permissions"},
			{Name: "system_user_add_group"},
			{Name: "system_user_delete_group"},
		}
		DB.Create(&p)
	}
	DB.Model(&Group{}).Count(&count)
	if count == 0 {
		p := []Group{
			{Name: "user"},
			{Name: "super_user"},
			{Name: "power_user"},
		}
		DB.Create(&p)
	}
}
