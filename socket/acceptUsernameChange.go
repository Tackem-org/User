package socket

import (
	_ "image/gif"
	_ "image/jpeg"
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	pb "github.com/Tackem-org/Proto/pb/web"
	"github.com/Tackem-org/User/model"
)

func AcceptUsernameChange(in *system.WebSocketRequest) (*system.WebSocketReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[socket.AcceptUsernameChange(in *system.WebSocketRequest) (*system.WebSocketReturn, error)]")
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

	result3 := model.DB.Model(&user).Update("Username", usernameRequest.Name)
	if result3.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "username rename failed possably already exists",
		}, nil
	}
	in.Data["name"] = usernameRequest.Name
	taskID := usernameRequest.ID
	result4 := model.DB.Delete(&usernameRequest)
	if result4.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "failed to delete the request",
		}, nil
	}
	system.RemoveTask(&pb.RemoveTaskRequest{
		Task:   "usernamechangerequest",
		BaseId: system.RegData().GetBaseID(),
		TaskId: taskID,
	})
	in.Data["updatedat"] = user.UpdatedAt
	return &system.WebSocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
