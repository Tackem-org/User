package editUser

import (
	"net/http"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

func DeleteUser(in *structs.SocketRequest) (*structs.SocketReturn, error) {
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

	model.DB.Delete(&user)
	in.Data["updatedat"] = user.UpdatedAt
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
