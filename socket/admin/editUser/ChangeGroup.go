package editUser

import (
	"net/http"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

func ChangeGroup(in *structs.SocketRequest) (*structs.SocketReturn, error) {
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

	tmpGroup, foundGroup := in.Data["group"]
	if !foundGroup {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "group missing",
		}, nil
	}
	valGroup, okGroup := tmpGroup.(int)
	if !okGroup {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "group not a int",
		}, nil
	}

	var group model.Group
	model.DB.Where(&model.Group{ID: uint64(valGroup)}).First(&group)
	if group.ID == 0 {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "group not found",
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
		model.DB.Model(&user).Association("Groups").Append(&group)
	} else {
		model.DB.Model(&user).Association("Groups").Delete(&group)
	}

	in.Data["updatedat"] = user.UpdatedAt
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
