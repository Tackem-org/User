package editUser

import (
	_ "image/gif"
	_ "image/jpeg"
	"net/http"
	"regexp"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

func ChangeUsername(in *system.WebSocketRequest) (*system.WebSocketReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[socket.admin.editUser.ChangeUsername(in *system.WebSocketRequest) (*system.WebSocketReturn, error)]")
	uidf, ok := in.Data["userid"].(float64)
	if !ok {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "userid not valid",
		}, nil
	}
	userID := uint64(uidf)
	var user model.User
	result := model.DB.Preload(clause.Associations).Find(&user, userID)
	if result.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "user not found",
		}, nil
	}
	val, ok := in.Data["username"].(string)
	if !ok || val == "" || len(val) <= 4 || !regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString(val) {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "username not valid",
		}, nil
	}
	result2 := model.DB.Model(&user).Update("Username", val)
	if result2.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "username already exists " + result2.Error.Error(),
		}, nil
	}

	in.Data["updatedat"] = user.UpdatedAt
	return &system.WebSocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
