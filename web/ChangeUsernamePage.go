package web

import (
	"net/http"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
)

func ChangeUsernamePage(in *structs.WebRequest) (*structs.WebReturn, error) {
	if !in.User.HasPermission("system_user_change_own_username") {
		return &structs.WebReturn{
			StatusCode:   http.StatusForbidden,
			ErrorMessage: "user not authorised to view this page",
		}, nil
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
