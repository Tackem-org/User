package web

import (
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
)

type arpud struct {
	ID       uint64
	Username string
	Password string
	Disabled bool
	IsAdmin  bool
}

func AdminRootPage(in *system.WebRequest) (*system.WebReturn, error) {
	user := model.User{}
	users := []arpud{}
	rows, _ := model.DB.Find(&user).Rows()
	for rows.Next() {
		// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
		model.DB.ScanRows(rows, &user)
		users = append(users, arpud{
			ID:       uint64(user.ID),
			Username: user.Username,
			Password: "",
			Disabled: user.Disabled,
			IsAdmin:  user.IsAdmin,
		})
	}
	return &system.WebReturn{
		FilePath: "admin/root",
		PageData: map[string]interface{}{
			"Users": users,
			"Test":  "Testing Admin Data Here",
		},
	}, nil
}
