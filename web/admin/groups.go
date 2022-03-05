package admin

import (
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/structs"
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

func AdminGroupsPage(in *structs.WebRequest) (*structs.WebReturn, error) {
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
			Title:       strings.ReplaceAll(group.Name, "_", " "),
			Permissions: groupPermissions,
			UserCount:   len(group.Users),
		})
	}

	return &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "admin/groups",
		PageData: map[string]interface{}{
			"Groups":       groupsList,
			"Permissions":  allPermissionsList,
			"CreateLength": len(allPermissionsList) + 3,
		},
	}, nil
}
