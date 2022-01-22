package web

import (
	"strings"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
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
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminRootPage(in *system.WebRequest) (*system.WebReturn, error)]")
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
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminUserIDPage(in *system.WebRequest) (*system.WebReturn, error)]")
	var user model.User
	model.DB.First(&user, in.PathVariables["userid"])
	return &system.WebReturn{
		FilePath: "admin/user",
		PageData: map[string]interface{}{
			"User": user,
		},
	}, nil
}

func AdminGroupsPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminGroupsPage(in *system.WebRequest) (*system.WebReturn, error)]")

	var allPermissions []model.Permission
	model.DB.Find(&allPermissions)
	var allPermissionsList []sPermissions
	for _, permission := range allPermissions {
		allPermissionsList = append(allPermissionsList, sPermissions{
			ID:    permission.ID,
			Name:  permission.Name,
			Title: strings.ReplaceAll(permission.Name, "_", " "),
		})
	}

	var groups []model.Group
	model.DB.Preload("Permissions").Find(&groups)
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
			"Groups":      groupsList,
			"Permissions": allPermissionsList,
		},
	}, nil
}

func AdminPermissionsPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminPermissionsPage(in *system.WebRequest) (*system.WebReturn, error)]")
	var p []model.Permission
	model.DB.Find(&p)
	var sp []sPermissions
	for _, r := range p {
		sp = append(sp, sPermissions{
			ID:          r.ID,
			Name:        r.Name,
			GroupsCount: len(r.Groups),
			UserCount:   len(r.Users),
		})
	}

	return &system.WebReturn{
		FilePath: "admin/permissions",
		PageData: map[string]interface{}{
			"permissions": sp,
		},
	}, nil
}

type sGroups struct {
	ID          uint64
	Name        string
	Permissions []sPermissions
	UserCount   int
	Active      bool
}

type sPermissions struct {
	ID          uint64
	Name        string
	Title       string
	GroupsCount int
	UserCount   int
	Active      bool
}

func checkActivePermissions(findID uint64, gp []model.Permission) bool {
	for _, enabledPermission := range gp {
		if findID == enabledPermission.ID {
			return true
		}
	}
	return false
}
