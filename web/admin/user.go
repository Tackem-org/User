package admin

import (
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

func AdminUserIDPage(in *structs.WebRequest) (*structs.WebReturn, error) {
	tmpUserID, foundUserID := in.PathVariables["userid"]
	if !foundUserID {
		return &structs.WebReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "userid missing",
		}, nil
	}
	userID, okUserID := tmpUserID.(int)
	if !okUserID {
		return &structs.WebReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "userid not an int",
		}, nil
	}
	var user model.User
	model.DB.Preload(clause.Associations).Where(&model.User{ID: uint64(userID)}).Find(&user)
	if user.ID == 0 {
		return &structs.WebReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "user not found",
		}, nil
	}

	var allPermissions []model.Permission
	var allPermissionsList []sPermissions
	var allGroups []model.Group
	var allGroupsList []sGroups
	var usernameRequest model.UsernameRequest

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

	model.DB.Where(&model.UsernameRequest{RequestUserID: user.ID}).First(&usernameRequest)
	var ur string
	var urid uint64
	if usernameRequest.ID == 0 {
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
