package admin

import (
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

type sPermissions struct {
	ID          uint64
	Name        string
	Title       string
	GroupsCount int
	UserCount   int
	Active      bool
}

func AdminPermissionsPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminPermissionsPage(in *system.WebRequest) (*system.WebReturn, error)]")
	var permissions []model.Permission
	model.DB.Preload(clause.Associations).Find(&permissions)
	var sortedPermissions []sPermissions
	for _, r := range permissions {
		sortedPermissions = append(sortedPermissions, sPermissions{
			ID:          r.ID,
			Name:        r.Name,
			Title:       strings.ReplaceAll(r.Name, "_", " "),
			GroupsCount: len(r.Groups),
			UserCount:   len(r.Users),
		})
	}

	return &system.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "admin/permissions",
		PageData: map[string]interface{}{
			"permissions": sortedPermissions,
		},
	}, nil
}

func checkActivePermissions(findID uint64, permissions []model.Permission) bool {
	for _, enabledPermission := range permissions {
		if findID == enabledPermission.ID {
			return true
		}
	}
	return false
}
