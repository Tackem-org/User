package admin

import (
	"net/http"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

type AdminUserData struct {
	ID               uint64
	Username         string
	Password         string
	Disabled         bool
	IsAdmin          bool
	GroupsCount      int
	PermissionsCount int
}

func AdminRootPage(in *structs.WebRequest) (*structs.WebReturn, error) {
	var users []model.User
	lusers := []AdminUserData{}
	model.DB.Preload(clause.Associations).Find(&users)
	for _, user := range users {
		lusers = append(lusers, AdminUserData{
			ID:               uint64(user.ID),
			Username:         user.Username,
			Password:         "",
			Disabled:         user.Disabled,
			IsAdmin:          user.IsAdmin,
			GroupsCount:      len(user.Groups),
			PermissionsCount: len(user.Permissions),
		})
	}
	return &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "admin/root",
		PageData: map[string]interface{}{
			"Users": lusers,
		},
	}, nil
}
