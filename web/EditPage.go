package web

import (
	"net/http"

	"github.com/Tackem-org/Global/structs"
)

func EditPage(in *structs.WebRequest) (*structs.WebReturn, error) {
	return &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "edit",
		PageData:   map[string]interface{}{},
	}, nil
}
