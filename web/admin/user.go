package admin

import (
	"strings"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

func AdminUserIDPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminUserIDPage(in *system.WebRequest) (*system.WebReturn, error)]")

	var user model.User
	model.DB.Preload(clause.Associations).First(&user, in.PathVariables["userid"])

	var allPermissions []model.Permission
	model.DB.Find(&allPermissions)
	var allPermissionsList []sPermissions
	for _, permission := range allPermissions {
		p := sPermissions{
			ID:    permission.ID,
			Name:  permission.Name,
			Title: strings.ReplaceAll(permission.Name, "_", " "),
		}
		for _, v := range user.Permissions {
			if v.Name == permission.Name {
				p.Active = true
				break
			}
		}
		allPermissionsList = append(allPermissionsList, p)
	}

	var allGroups []model.Group
	model.DB.Find(&allGroups)
	var allGroupsList []sGroups
	for _, group := range allGroups {
		g := sGroups{
			ID:    group.ID,
			Name:  group.Name,
			Title: strings.ReplaceAll(group.Name, "_", " "),
		}
		for _, v := range user.Groups {
			if v.Name == group.Name {
				g.Active = true
				break
			}
		}
		allGroupsList = append(allGroupsList, g)
	}

	return &system.WebReturn{
		FilePath:       "admin/user",
		CustomPageName: "admin-user-edit",
		PageData: map[string]interface{}{
			"User":        user,
			"Permissions": allPermissionsList,
			"Groups":      allGroupsList,
		},
	}, nil
}

func AdminEditUserWebSocket(in *system.WebSocketRequest) (*system.WebSocketReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminEditUserWebSocket(in *system.WebSocketRequest) (*system.WebSocketReturn, error)]")

	d := in.Data
	// userID := d["userid"]
	switch d["command"] {
	// case "changeusername":
	// case "updatepassword":
	// case "changedisabled":
	// case "changeisadmin":
	// case "deleteuser":
	// case "setgroup":
	// case "setpermission":

	default:
		return &system.WebSocketReturn{
			StatusCode:   200,
			ErrorMessage: "command not found",
			Data:         map[string]interface{}{},
		}, nil
	}
}
