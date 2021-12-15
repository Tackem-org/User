package database

import (
	"github.com/Tackem-org/User/flags"
	"github.com/Tackem-org/User/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Setup() {

	DB, err := gorm.Open(sqlite.Open(*flags.DatabaseFile), &gorm.Config{})
	if err != nil {
		panic("failed to Open database")
	}

	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Group{})

}
