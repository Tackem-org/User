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
		},
	}, nil
}

func AdminUserIDPage(in *system.WebRequest) (*system.WebReturn, error) {
	var user model.User
	model.DB.First(&user, in.PathVariables["userid"])
	return &system.WebReturn{
		FilePath: "admin/user",
		PageData: map[string]interface{}{
			"User": user,
		},
	}, nil
}
