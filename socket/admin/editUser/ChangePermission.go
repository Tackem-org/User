package editUser

import (
	"net/http"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

func ChangePermission(in *structs.SocketRequest) (*structs.SocketReturn, error) {
	tmpUserID, foundUserID := in.Data["userid"]
	if !foundUserID {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "userid missing",
		}, nil
	}
	userID, okUserID := tmpUserID.(int)
	if !okUserID {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "userid not an int",
		}, nil
	}
	var user model.User
	model.DB.Preload(clause.Associations).Where(&model.User{ID: uint64(userID)}).Find(&user)
	if user.ID == 0 {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "user not found",
		}, nil
	}

	tmpPermission, foundPermission := in.Data["permission"]
	if !foundPermission {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "permission missing",
		}, nil
	}
	valPermission, okPermission := tmpPermission.(int)
	if !okPermission {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "permission not a int",
		}, nil
	}

	var permission model.Permission
	model.DB.Where(&model.Permission{ID: uint64(valPermission)}).First(&permission)
	if permission.ID == 0 {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "permission not found",
		}, nil
	}

	tmpChecked, foundChecked := in.Data["checked"]
	if !foundChecked {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "checked missing",
		}, nil
	}
	valChecked, okChecked := tmpChecked.(bool)
	if !okChecked {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "checked not a bool",
		}, nil
	}

	if valChecked == true {
		model.DB.Model(&user).Association("Permissions").Append(&permission)
	} else {
		model.DB.Model(&user).Association("Permissions").Delete(&permission)
	}

	in.Data["updatedat"] = user.UpdatedAt
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
