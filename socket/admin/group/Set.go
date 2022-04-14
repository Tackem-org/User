package group

import (
	"net/http"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
)

func Set(in *structs.SocketRequest) (*structs.SocketReturn, error) {
	tmpGroupid, foundgid := in.Data["groupid"]
	if !foundgid {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "GroupID Missing",
		}, nil
	}
	tmpPermissionid, foundpid := in.Data["permissionid"]
	if !foundpid {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "PermissionID Missing",
		}, nil
	}
	tmpChecked, foundchecked := in.Data["checked"]
	if !foundchecked {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "checked Missing",
		}, nil
	}
	groupid := uint64(tmpGroupid.(int))
	permissionid := uint64(tmpPermissionid.(int))
	checked := tmpChecked.(bool)

	var group model.Group
	var permission model.Permission
	model.DB.Where(&model.Group{ID: groupid}).First(&group)
	if group.ID == 0 {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "Group Not Found",
		}, nil
	}

	model.DB.Where(&model.Permission{ID: permissionid}).First(&permission, permissionid)
	if permission.ID == 0 {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "Permission Not Found",
		}, nil
	}

	if checked {
		model.DB.Model(&group).Association("Permissions").Append(&permission)
	} else {
		model.DB.Model(&group).Association("Permissions").Delete(&permission)
	}

	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
