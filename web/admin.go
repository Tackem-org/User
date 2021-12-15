package web

import (
	"github.com/Tackem-org/Global/remoteWebSystem"
)

func AdminRootPage(in *remoteWebSystem.WebRequest) (*remoteWebSystem.WebReturn, error) {
	return &remoteWebSystem.WebReturn{
		FilePath: "admin/root",
		PageData: map[string]interface{}{
			"Test": "Testing Admin Data Here",
		},
	}, nil
}
