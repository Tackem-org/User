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
	var p []model.Permission
	model.DB.Preload(clause.Associations).Find(&p)
	var sp []sPermissions
	for _, r := range p {
		sp = append(sp, sPermissions{
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
			"permissions": sp,
		},
	}, nil
}

func checkActivePermissions(findID uint64, gp []model.Permission) bool {
	for _, enabledPermission := range gp {
		if findID == enabledPermission.ID {
			return true
		}
	}
	return false
}
