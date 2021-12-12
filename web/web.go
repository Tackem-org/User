package web

import (
	"github.com/Tackem-org/Global/remoteWebSystem"
	"github.com/Tackem-org/User/static"
)

func Setup() {
	remoteWebSystem.Setup(&static.FS)
	remoteWebSystem.AddPath("/", RootPage)
	remoteWebSystem.AddPath("admin/", AdminRootPage)
}

func RootPage(in *remoteWebSystem.WebRequest) *remoteWebSystem.WebReturn {
	return &remoteWebSystem.WebReturn{
		FilePath: "root",
		PageData: map[string]interface{}{
			"Test": "Testing Data Here",
		},
	}
}

func AdminRootPage(in *remoteWebSystem.WebRequest) *remoteWebSystem.WebReturn {
	return &remoteWebSystem.WebReturn{
		FilePath: "admin/root",
		PageData: map[string]interface{}{
			"Test": "Testing Data Here",
		},
	}
}
