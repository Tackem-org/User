package admin

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AdminUserIDPage(in *structs.WebRequest) (*structs.WebReturn, error) {
	var userID uint64
	useridvar, found := in.PathVariables["userid"]
	if !found {
		return &structs.WebReturn{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "userid not found from path",
		}, nil
	}
	userID = uint64(useridvar.(float64))
	var user model.User
	var allPermissions []model.Permission
	var allPermissionsList []sPermissions
	var allGroups []model.Group
	var allGroupsList []sGroups
	var usernameRequest model.UsernameRequest

	model.DB.Preload(clause.Associations).First(&user, userID)
	model.DB.Find(&allPermissions)
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

	model.DB.Find(&allGroups)
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

	result := model.DB.Where(&model.UsernameRequest{RequestUserID: userID}).First(&usernameRequest)
	var ur string
	var urid uint64
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ur = usernameRequest.Name
		urid = usernameRequest.ID
	}

	return &structs.WebReturn{
		StatusCode:     http.StatusOK,
		FilePath:       "admin/user",
		CustomPageName: "admin-user-edit",
		PageData: map[string]interface{}{
			"User":              user,
			"Permissions":       allPermissionsList,
			"Groups":            allGroupsList,
			"UsernameRequest":   ur,
			"UsernameRequestID": urid,
		},
	}, nil
}
