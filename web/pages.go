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
		PageData: map[string]interface{}{
			"Test": "Testing Data Here",
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

func UserIDPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.UserIDPage(in *system.WebRequest) (*system.WebReturn, error)]")
	return &system.WebReturn{
		FilePath: "view/userid",
		PageData: map[string]interface{}{
			"UserID": in.PathVariables["userid"],
		},
	}, nil
}

func UserNamePage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.UserNamePage(in *system.WebRequest) (*system.WebReturn, error)]")
	return &system.WebReturn{
		FilePath: "view/username",
		PageData: map[string]interface{}{
			"UserName": in.PathVariables["username"],
		},
	}, nil
}
