package admin

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

type UserData struct {
	ID               uint64
	Username         string
	Password         string
	Disabled         bool
	IsAdmin          bool
	GroupsCount      int
	PermissionsCount int
}

func AdminRootPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminRootPage(in *system.WebRequest) (*system.WebReturn, error)]")
	var users []model.User
	lusers := []UserData{}
	model.DB.Preload(clause.Associations).Find(&users)
	for _, user := range users {
		lusers = append(lusers, UserData{
			ID:               uint64(user.ID),
			Username:         user.Username,
			Password:         "",
			Disabled:         user.Disabled,
			IsAdmin:          user.IsAdmin,
			GroupsCount:      len(user.Groups),
			PermissionsCount: len(user.Permissions),
		})
	}
	return &system.WebReturn{
		FilePath: "admin/root",
		PageData: map[string]interface{}{
			"Users": lusers,
		},
	}, nil
}
