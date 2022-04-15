package editUser

import (
	"net/http"

	"github.com/Tackem-org/Global/config"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
	"gorm.io/gorm/clause"
)

func ChangePassword(in *structs.SocketRequest) (*structs.SocketReturn, error) {
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

	tmpPassword, foundPassword := in.Data["password"]
	if !foundPassword {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "password missing",
		}, nil
	}
	val, ok := tmpPassword.(string)
	if !ok {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "password not a string",
		}, nil
	}

	minPassLength, _ := config.GetUint("user.password.minimum")
	if uint(len(val)) <= minPassLength {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "password too short",
		}, nil
	}
	model.DB.Model(&user).Update("Password", password.Hash(val))
	in.Data["updatedat"] = user.UpdatedAt
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
