package web

import (
	"net/http"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/web"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
	"github.com/Tackem-org/User/tasks"
)

func RequestUsernamePage(in *structs.WebRequest) (*structs.WebReturn, error) {
	var user model.User
	var usernameRequest model.UsernameRequest

	success := false
	errorString := ""

	model.DB.Where(&model.UsernameRequest{RequestUserID: in.User.ID}).Find(&usernameRequest)
	if usernameRequest.RequestUserID == in.User.ID {
		return &structs.WebReturn{
			StatusCode: http.StatusOK,
			FilePath:   "requestusername",
			PageData: map[string]interface{}{
				"RequestMade": true,
				"Success":     false,
				"Error":       "",
				"Username":    in.User.Name,
			},
		}, nil
	}

	if len(in.Post) > 0 {
		username, ok1 := in.Post["username"].(string)
		pword, ok2 := in.Post["password"].(string)
		if !ok1 || !ok2 {
			errorString = "error cannot get post data"
		} else if username == "" {
			errorString = "new username cannot be blank"
		} else if username == in.User.Name {
			errorString = "username same as previous"
		} else if pword == "" {
			errorString = "password cannot be blank"
		} else {
			var existing model.User
			model.DB.Where(&model.User{Username: username}).First(&existing)
			if existing.ID > 0 {
				errorString = "user already exists"
			} else {
				model.DB.Where(&model.User{ID: in.User.ID, Password: password.Hash(pword)}).First(&user)
				if user.ID != in.User.ID {
					errorString = "password doesn't match"
				} else {
					usernameRequest.Name = username
					usernameRequest.RequestUserID = in.User.ID
					model.DB.Create(&usernameRequest)
					success = true
					usernameRequest.RequestUser = user
					web.AddTask(tasks.UserNameChangeRequest(&usernameRequest))
				}
			}
		}
	}

	return &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "requestusername",
		PageData: map[string]interface{}{
			"RequestMade": false,
			"Success":     success,
			"Error":       errorString,
			"Username":    in.User.Name,
		},
	}, nil
}
