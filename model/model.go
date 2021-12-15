package model

import (
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique"`
	Password    string
	IsAdmin     bool
	Groups      []*Group      `gorm:"many2many:groups;"`
	Permissions []*Permission `gorm:"many2many:permissions;"`
}

type Group struct {
	gorm.Model
	Name        string
	Users       []*User       `gorm:"many2many:users;"`
	SubGroups   []*Group      `gorm:"many2many:sub_groups;"`
	Permissions []*Permission `gorm:"many2many:permissions;"`
}

type Permission struct {
	gorm.Model
	Name   string
	Users  []*User  `gorm:"many2many:users;"`
	Groups []*Group `gorm:"many2many:groups;"`
}

func hashPassword(password string) string {
	// f, err := os.Open("/salt.dat")
	// defer f.Close()
	// if err != nil {

	// }
	//TODO grab the salt from the config

	salt := make([]byte, 12)

	return fmt.Sprintf("%x", pbkdf2.Key([]byte(password), salt, 4096, len(salt)*8, sha512.New))
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

}
