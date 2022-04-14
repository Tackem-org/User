package socket

import (
	"net/http"

	pb "github.com/Tackem-org/Global/pb/web"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/web"
	"github.com/Tackem-org/User/model"
)

func AcceptUsernameChange(in *structs.SocketRequest) (*structs.SocketReturn, error) {
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

	var user2 model.User
	model.DB.Where(&model.User{Username: usernameRequest.Name}).First(&user2)
	if user2.ID > 0 {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "username rename failed already exists",
		}, nil
	}

	model.DB.Model(&user).Update("Username", usernameRequest.Name)
	in.Data["name"] = usernameRequest.Name
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
