package socket

import (
	_ "image/gif"
	_ "image/jpeg"
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
)

func RejectUsernameChange(in *system.WebSocketRequest) (*system.WebSocketReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[socket.RejectUsernameChange(in *system.WebSocketRequest) (*system.WebSocketReturn, error)]")
	val, ok := in.Data["userid"].(float64)
	if !ok {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "userid not valid",
		}, nil
	}
	userID := uint64(val)
	var usernameRequest model.UsernameRequest
	var user model.User
	result1 := model.DB.Find(&user, userID)
	result2 := model.DB.Where(&model.UsernameRequest{RequestUserID: uint64(userID)}).First(&usernameRequest)
	if result1.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "userid not found",
		}, nil
	}
	if result2.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "username request not found",
		}, nil
	}

	result4 := model.DB.Delete(&usernameRequest)
	if result4.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "failed to delete the request",
		}, nil
	}

	in.Data["updatedat"] = user.UpdatedAt
	return &system.WebSocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
