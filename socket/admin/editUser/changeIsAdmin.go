package editUser

import (
	_ "image/gif"
	_ "image/jpeg"
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

func ChangeIsAdmin(in *system.WebSocketRequest) (*system.WebSocketReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[socket.admin.editUser.ChangeIsAdmin(in *system.WebSocketRequest) (*system.WebSocketReturn, error)]")
	userID := in.Data["userid"]
	var user model.User
	result := model.DB.Preload(clause.Associations).Find(&user, userID)
	if result.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "user not found",
		}, nil
	}
	val, ok := in.Data["checked"].(bool)
	if !ok {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "changing disabled failed",
		}, nil
	}
	result2 := model.DB.Model(&user).Update("IsAdmin", val)
	if result2.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "isadmin error " + result2.Error.Error(),
		}, nil
	}
	in.Data["updatedat"] = user.UpdatedAt
	return &system.WebSocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
