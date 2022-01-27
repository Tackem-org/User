package web

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
)

func RootPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.RootPage(in *system.WebRequest) (*system.WebReturn, error)]")
	return &system.WebReturn{
		FilePath: "root",
		PageData: map[string]interface{}{},
	}, nil
}

func PasswordPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.PasswordPage(in *system.WebRequest) (*system.WebReturn, error)]")
	success := false
	errorString := ""
	if len(in.Post) > 0 {
		//TODO MAKE THIS ACTUALLY CHECK THEN UPDATE THE PASSWORD IN DB
		logging.Infof("POST= %+v", in.Post)
		success = true
	}
	return &system.WebReturn{
		FilePath: "password",
		PageData: map[string]interface{}{
			"Success": success,
			"Error":   errorString,
		},
	}, nil
}

func EditPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.EditPage(in *system.WebRequest) (*system.WebReturn, error)]")
	return &system.WebReturn{
		FilePath: "edit",
		PageData: map[string]interface{}{},
	}, nil
}
