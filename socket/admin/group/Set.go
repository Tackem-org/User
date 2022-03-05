package group

import (
	"errors"
	"net/http"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm"
)

func Set(in *structs.SocketRequest) (*structs.SocketReturn, error) {
	var group model.Group
	var permission model.Permission
	fgroupid, okgid := in.Data["groupid"].(float64)
	fpermissionid, okpid := in.Data["permissionid"].(float64)
	if !okgid || !okpid {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "ids Not Numbers",
		}, nil
	}
	groupid := uint64(fgroupid)
	permissionid := uint64(fpermissionid)
	if err := model.DB.First(&group, groupid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &structs.SocketReturn{
				StatusCode:   http.StatusNotFound,
				ErrorMessage: "Group Not Found",
			}, nil
		}
		return &structs.SocketReturn{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "DB Group ERROR: " + err.Error(),
		}, nil
	}
	if err := model.DB.First(&permission, permissionid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &structs.SocketReturn{
				StatusCode:   http.StatusNotFound,
				ErrorMessage: "Permission Not Found",
			}, nil
		}
		return &structs.SocketReturn{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "DB Group ERROR: " + err.Error(),
		}, nil
	}

	if in.Data["checked"] == true {
		if err := model.DB.Model(&group).Association("Permissions").Append(&permission); err != nil {
			return &structs.SocketReturn{
				StatusCode:   http.StatusInternalServerError,
				ErrorMessage: "DB Permission Append ERROR: " + err.Error(),
			}, nil
		}

	} else {
		if err := model.DB.Model(&group).Association("Permissions").Delete(&permission); err != nil {
			return &structs.SocketReturn{
				StatusCode:   http.StatusInternalServerError,
				ErrorMessage: "DB Permission Delete ERROR: " + err.Error(),
			}, nil
		}
	}

	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
