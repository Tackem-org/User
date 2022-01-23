package admin

import (
	"strings"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

type sGroups struct {
	ID          uint64
	Name        string
	Title       string
	Permissions []sPermissions
	UserCount   int
	Active      bool
}

func AdminGroupsPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminGroupsPage(in *system.WebRequest) (*system.WebReturn, error)]")

	var allPermissions []model.Permission
	model.DB.Preload(clause.Associations).Find(&allPermissions)
	var allPermissionsList []sPermissions
	for _, permission := range allPermissions {
		allPermissionsList = append(allPermissionsList, sPermissions{
			ID:    permission.ID,
			Name:  permission.Name,
			Title: strings.ReplaceAll(permission.Name, "_", " "),
		})
	}

	var groups []model.Group
	model.DB.Preload(clause.Associations).Find(&groups)
	var groupsList []sGroups
	for _, group := range groups {
		var groupPermissions []sPermissions

		for _, permission := range allPermissionsList {
			permission.Active = checkActivePermissions(permission.ID, group.Permissions)
			groupPermissions = append(groupPermissions, permission)
		}
		groupsList = append(groupsList, sGroups{
			ID:          group.ID,
			Name:        group.Name,
			Permissions: groupPermissions,
			UserCount:   len(group.Users),
		})
	}

	return &system.WebReturn{
		FilePath: "admin/groups",
		PageData: map[string]interface{}{
			"Groups":       groupsList,
			"Permissions":  allPermissionsList,
			"CreateLength": len(allPermissionsList) + 3,
		},
	}, nil
}

func AdminGroupsWebSocket(in *system.WebSocketRequest) (*system.WebSocketReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminGroupsWebSocket(in *system.WebSocketRequest) (*system.WebSocketReturn, error)]")

	d := in.Data
	switch d["command"] {
	case "setgroup":
		var g model.Group
		var p model.Permission
		model.DB.First(&g, d["groupid"])
		model.DB.First(&p, d["permissionid"])
		if d["checked"] == true {
			model.DB.Model(&g).Association("Permissions").Append(&p)
			model.DB.Save(&g)
		} else {
			model.DB.Model(&g).Association("Permissions").Delete(&p)
			model.DB.Save(&g)
		}
		return &system.WebSocketReturn{
			StatusCode:   200,
			ErrorMessage: "",
			Data:         d,
		}, nil
	case "addgroup":
		g := model.Group{
			Name: d["name"].(string),
		}
		result := model.DB.Create(&g)
		if result.Error != nil {
			return &system.WebSocketReturn{
				StatusCode:   400,
				ErrorMessage: "New Group Name Must Be Unique",
				Data:         d,
			}, nil
		}
		d["groupid"] = g.ID
		return &system.WebSocketReturn{
			StatusCode:   200,
			ErrorMessage: "",
			Data:         d,
		}, nil
	case "deletegroup":
		var g model.Group
		model.DB.First(&g, d["groupid"])
		model.DB.Delete(&g)
		return &system.WebSocketReturn{
			StatusCode:   200,
			ErrorMessage: "",
			Data:         d,
		}, nil
	default:
		return &system.WebSocketReturn{
			StatusCode:   200,
			ErrorMessage: "command not found",
			Data:         map[string]interface{}{},
		}, nil
	}
}
