package editUser

import (
	"net/http"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

func ChangeIsAdmin(in *structs.SocketRequest) (*structs.SocketReturn, error) {
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
	tmpChecked, foundChecked := in.Data["checked"]
	if !foundChecked {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "checked missing",
		}, nil
	}
	val, ok := tmpChecked.(bool)
	if !ok {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "checked not a bool",
		}, nil
	}
	model.DB.Model(&user).Update("IsAdmin", val)
	in.Data["updatedat"] = user.UpdatedAt
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
