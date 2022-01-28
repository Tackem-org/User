package web

import (
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
)

func RootPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.RootPage(in *system.WebRequest) (*system.WebReturn, error)]")
	return &system.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "root",
		PageData:   map[string]interface{}{},
	}, nil
}

func PasswordPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.PasswordPage(in *system.WebRequest) (*system.WebReturn, error)]")

	if !in.User.HasPermission("system_user_change_own_password") {
		return &system.WebReturn{
			StatusCode:   http.StatusForbidden,
			ErrorMessage: "user not authorised to view this page",
		}, nil
	}
	success := false
	errorString := ""
	if len(in.Post) > 0 {
		op, ok1 := in.Post["op"].(string)
		np1, ok2 := in.Post["np1"].(string)
		np2, ok3 := in.Post["np2"].(string)

		if !ok1 || !ok2 || !ok3 {
			logging.Infof("op:%t, np1:%t, np2:%t", ok1, ok2, ok3)
		}
		logging.Infof("op:%s, np1:%s, np2:%s", op, np1, np2)
		if op == "" {
			errorString = "original password blank"
		} else if np1 == "" || np2 == "" {
			errorString = "new password cannot be blank"
		} else if len(np1) <= 4 || len(np2) <= 4 {
			errorString = "new password too short"
		} else if np1 != np2 {
			errorString = "new password don't match"
		} else {
			var user model.User
			model.DB.Where(&model.User{ID: in.User.ID, Password: password.Hash(op)}).First(&user)
			if user.ID != in.User.ID {
				errorString = "old password doesn't match"
			} else {
				user.Password = password.Hash(np1)
				model.DB.Save(&user)
				success = true
			}
		}
	}
	return &system.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "password",
		PageData: map[string]interface{}{
			"Success": success,
			"Error":   errorString,
		},
	}, nil
}

func EditPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.EditPage(in *system.WebRequest) (*system.WebReturn, error)]")
	return &system.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "edit",
		PageData:   map[string]interface{}{},
	}, nil
}
