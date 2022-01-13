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

func UserNamePage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.UserNamePage(in *system.WebRequest) (*system.WebReturn, error)]")
	if in.PathVariables["username"] == "Tom" {
		return &system.WebReturn{
			FilePath: "user",
			PageData: map[string]interface{}{
				"Test":   "Good User Name",
				"userid": in.PathVariables["username"],
			},
		}, nil
	}
	return &system.WebReturn{
		FilePath: "user",
		PageData: map[string]interface{}{
			"Test":   "Bad User Name",
			"userid": in.PathVariables["username"],
		},
	}, nil
}

func UserIDPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.UserIDPage(in *system.WebRequest) (*system.WebReturn, error)]")
	if in.PathVariables["userid"] == "1" {
		return &system.WebReturn{
			FilePath: "user",
			PageData: map[string]interface{}{
				"Test":   "Good User ID",
				"userid": in.PathVariables["userid"],
			},
		}, nil
	}
	return &system.WebReturn{
		FilePath: "user",
		PageData: map[string]interface{}{
			"Test":   "Bad user ID",
			"userid": in.PathVariables["userid"],
		},
	}, nil
}
