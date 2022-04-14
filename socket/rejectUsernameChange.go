package socket

import (
	"net/http"

	pb "github.com/Tackem-org/Global/pb/web"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/web"
	"github.com/Tackem-org/User/model"
)

func RejectUsernameChange(in *structs.SocketRequest) (*structs.SocketReturn, error) {
	val, ok := in.Data["userid"].(int)
	if !ok {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "userid not valid",
		}, nil
	}
	userID := uint64(val)
	var usernameRequest model.UsernameRequest
	var user model.User
	result1 := model.DB.Where(&model.UsernameRequest{ID: userID}).First(&user)
	if result1.Error != nil {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "userid not found",
		}, nil
	}
	result2 := model.DB.Where(&model.UsernameRequest{RequestUserID: userID}).First(&usernameRequest)
	if result2.Error != nil {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "username request not found",
		}, nil
	}
	taskID := usernameRequest.ID
	model.DB.Delete(&usernameRequest)
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
