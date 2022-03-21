package socket

import (
	_ "image/gif"
	_ "image/jpeg"
	"net/http"

	pb "github.com/Tackem-org/Global/pb/web"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/web"
	"github.com/Tackem-org/User/model"
)

func RejectUsernameChange(in *structs.SocketRequest) (*structs.SocketReturn, error) {
	val, ok := in.Data["userid"].(float64)
	if !ok {
		return &structs.SocketReturn{
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
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "userid not found",
		}, nil
	}
	if result2.Error != nil {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "username request not found",
		}, nil
	}
	taskID := usernameRequest.ID
	result4 := model.DB.Delete(&usernameRequest)
	if result4.Error != nil {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "failed to delete the request",
		}, nil
	}
	web.RemoveTask(&pb.RemoveTaskRequest{
		Task:   "usernamechangerequest",
		TaskId: taskID,
	})
	in.Data["updatedat"] = user.UpdatedAt
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
