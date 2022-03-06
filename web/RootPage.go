package web

import (
	"net/http"

	"github.com/Tackem-org/Global/structs"
)

func RootPage(in *structs.WebRequest) (*structs.WebReturn, error) {
	return &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "root",
		PageData:   map[string]interface{}{},
	}, nil
}
