package web

import (
	"github.com/Tackem-org/Global/remoteWebSystem"
	"github.com/Tackem-org/User/static"
)

func Setup() {
	remoteWebSystem.Setup(&static.FS)
	remoteWebSystem.AddPath("/", RootPage)
	remoteWebSystem.AddAdminPath("/", AdminRootPage)
	remoteWebSystem.AddPath("{{number:userid}}", UserIDPage)
	remoteWebSystem.AddPath("{{string:username}}", UserNamePage)
}

func RootPage(in *remoteWebSystem.WebRequest) (*remoteWebSystem.WebReturn, error) {
	return &remoteWebSystem.WebReturn{
		FilePath: "root",
		PageData: map[string]interface{}{
			"Test": "Testing Data Here",
		},
	}, nil
}

func AdminRootPage(in *remoteWebSystem.WebRequest) (*remoteWebSystem.WebReturn, error) {
	return &remoteWebSystem.WebReturn{
		FilePath: "admin/root",
		PageData: map[string]interface{}{
			"Test": "Testing Admin Data Here",
		},
	}, nil
}

func UserNamePage(in *remoteWebSystem.WebRequest) (*remoteWebSystem.WebReturn, error) {
	if in.PathVariables["username"] == "Tom" {
		return &remoteWebSystem.WebReturn{
			FilePath: "user",
			PageData: map[string]interface{}{
				"Test":   "Good User Name",
				"userid": in.PathVariables["username"],
			},
		}, nil
	}
	return &remoteWebSystem.WebReturn{
		FilePath: "user",
		PageData: map[string]interface{}{
			"Test":   "Bad User Name",
			"userid": in.PathVariables["username"],
		},
	}, nil
}

func UserIDPage(in *remoteWebSystem.WebRequest) (*remoteWebSystem.WebReturn, error) {
	if in.PathVariables["userid"] == "1" {
		return &remoteWebSystem.WebReturn{
			FilePath: "user",
			PageData: map[string]interface{}{
				"Test":   "Good User ID",
				"userid": in.PathVariables["userid"],
			},
		}, nil
	}
	return &remoteWebSystem.WebReturn{
		FilePath: "user",
		PageData: map[string]interface{}{
			"Test":   "Bad user ID",
			"userid": in.PathVariables["userid"],
		},
	}, nil
}
