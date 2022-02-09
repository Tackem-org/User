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

func ChangeGroup(in *system.WebSocketRequest) (*system.WebSocketReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[socket.admin.editUser.ChangeGroup(in *system.WebSocketRequest) (*system.WebSocketReturn, error)]")
	userID := in.Data["userid"]
	var user model.User
	result := model.DB.Preload(clause.Associations).Find(&user, userID)
	if result.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "user not found",
		}, nil
	}

	var group model.Group
	result2 := model.DB.First(&group, in.Data["group"])
	if result2.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "group not found",
		}, nil
	}
	if in.Data["checked"] == true {
		model.DB.Model(&user).Association("Groups").Append(&group)
	} else {
		model.DB.Model(&user).Association("Groups").Delete(&group)
	}

	in.Data["updatedat"] = user.UpdatedAt
	return &system.WebSocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
