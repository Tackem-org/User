package web

import (
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
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
		//TODO MAKE THIS ACTUALLY CHECK THEN UPDATE THE PASSWORD IN DB
		logging.Infof("POST= %+v", in.Post)
		success = true
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
