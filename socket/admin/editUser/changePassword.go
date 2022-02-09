package editUser

import (
	_ "image/gif"
	_ "image/jpeg"
	"net/http"

	"github.com/Tackem-org/Global/config"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
	"gorm.io/gorm/clause"
)

func ChangePassword(in *system.WebSocketRequest) (*system.WebSocketReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[socket.admin.editUser.ChangePassword(in *system.WebSocketRequest) (*system.WebSocketReturn, error)]")
	userID := in.Data["userid"]
	var user model.User
	result := model.DB.Preload(clause.Associations).Find(&user, userID)
	if result.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "user not found",
		}, nil
	}

	val, ok := in.Data["password"].(string)
	if !ok {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "password not valid",
		}, nil
	}
	minPassLength, _ := config.GetUint("user.password.minimum")
	if uint(len(val)) <= minPassLength {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "password too short",
		}, nil
	}
	result2 := model.DB.Model(&user).Update("Password", password.Hash(val))
	if result2.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "password Error " + result2.Error.Error(),
		}, nil
	}

	in.Data["updatedat"] = user.UpdatedAt
	return &system.WebSocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
