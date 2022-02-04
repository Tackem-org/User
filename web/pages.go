package web

import (
	"net/http"

	"github.com/Tackem-org/Global/config"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
)

func RootPage(in *structs.WebRequest) (*structs.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.RootPage(in *structs.WebRequest) (*structs.WebReturn, error)]")
	return &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "root",
		PageData:   map[string]interface{}{},
	}, nil
}

func ChangePasswordPage(in *structs.WebRequest) (*structs.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.ChangePasswordPage(in *structs.WebRequest) (*structs.WebReturn, error)]")

	if !in.User.HasPermission("system_user_change_own_password") {
		return system.ForbiddenWebReturn()
	}
	minPassLength, _ := config.GetUint("user.password.minimum")
	success := false
	errorString := ""
	if len(in.Post) > 0 {
		op, ok1 := in.Post["op"].(string)
		np1, ok2 := in.Post["np1"].(string)
		np2, ok3 := in.Post["np2"].(string)

		if !ok1 || !ok2 || !ok3 {
			errorString = "error cannot get post data"
		} else if op == "" {
			errorString = "original password blank"
		} else if np1 == "" || np2 == "" {
			errorString = "new password cannot be blank"
		} else if uint(len(np1)) <= minPassLength || uint(len(np2)) <= minPassLength {
			errorString = "new password too short"
		} else if np1 != np2 {
			errorString = "new password don't match"
		} else {
			var user model.User
			model.DB.Where(&model.User{ID: in.User.ID, Password: password.Hash(op)}).First(&user)
			if user.ID != in.User.ID {
				errorString = "old password doesn't match"
			} else {
				result := model.DB.Model(&user).Update("Password", password.Hash(np1))
				if result.Error != nil {
					success = false
				} else {
					success = true
				}
			}
		}
	}
	return &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "changepassword",
		PageData: map[string]interface{}{
			"Success": success,
			"Error":   errorString,
		},
	}, nil
}

func ChangeUsernamePage(in *structs.WebRequest) (*structs.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.ChangeUsernamePage(in *structs.WebRequest) (*structs.WebReturn, error)]")
	if !in.User.HasPermission("system_user_change_own_username") {
		return system.ForbiddenWebReturn()
	}
	success := false
	errorString := ""
	if len(in.Post) > 0 {
		username, ok1 := in.Post["username"].(string)
		pword, ok2 := in.Post["password"].(string)
		if !ok1 || !ok2 {
			errorString = "error cannot get post data"
		} else if username == "" {
			errorString = "new username cannot be blank"
		} else if pword == "" {
			errorString = "password cannot be blank"
		} else {
			var user model.User
			model.DB.Where(&model.User{ID: in.User.ID, Password: password.Hash(pword)}).First(&user)
			if user.ID != in.User.ID {
				errorString = "old password doesn't match"
			} else {
				result := model.DB.Model(&user).Update("Username", username)
				if result.Error != nil {
					success = false
				} else {
					success = true
				}
			}
		}
	}
	return &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "changeusername",
		PageData: map[string]interface{}{
			"Success":  success,
			"Error":    errorString,
			"Username": in.User.Name,
		},
	}, nil
}

func RequestUsernamePage(in *structs.WebRequest) (*structs.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.RequestUsernamePage(in *structs.WebRequest) (*structs.WebReturn, error)]")
	if !in.User.HasPermission("system_user_request_change_of_username") {
		return system.ForbiddenWebReturn()
	}
	var user model.User
	var usernameRequest model.UsernameRequest

	requestMade := false
	success := false
	errorString := ""

	model.DB.Where(&model.UsernameRequest{RequestUserID: in.User.ID}).Find(&usernameRequest)
	if usernameRequest.RequestUserID == in.User.ID {
		requestMade = true
	} else {
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
			} else if err := model.DB.Where(&model.User{Username: username}).First(&user).Error; err == nil {
				errorString = "username already taken"
			} else {
				model.DB.Where(&model.User{ID: in.User.ID, Password: password.Hash(pword)}).First(&user)
				if user.ID != in.User.ID {
					errorString = "old password doesn't match"
				} else {
					usernameRequest.Name = username
					usernameRequest.RequestUserID = in.User.ID
					result := model.DB.Create(&usernameRequest)
					if result.Error != nil {
						success = false
					} else {
						success = true
					}
				}
			}
		}
	}

	return &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "requestusername",
		PageData: map[string]interface{}{
			"RequestMade": requestMade,
			"Success":     success,
			"Error":       errorString,
			"Username":    in.User.Name,
		},
	}, nil
}

func EditPage(in *structs.WebRequest) (*structs.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.EditPage(in *structs.WebRequest) (*structs.WebReturn, error)]")
	return &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "edit",
		PageData:   map[string]interface{}{},
	}, nil
}
