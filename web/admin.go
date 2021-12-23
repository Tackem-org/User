package web

import (
	"github.com/Tackem-org/Global/system"
)

func AdminRootPage(in *system.WebRequest) (*system.WebReturn, error) {
	return &system.WebReturn{
		FilePath: "admin/root",
		PageData: map[string]interface{}{
			"Test": "Testing Admin Data Here",
		},
	}, nil
}
