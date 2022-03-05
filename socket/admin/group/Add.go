package group

import (
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
)

func Add(in *structs.SocketRequest) (*structs.SocketReturn, error) {
	name := in.Data["name"].(string)
	if name == "" {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "New Group Name Cannot Be Blank",
		}, nil
	}
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ToLower(name)
	group := model.Group{
		Name: name,
	}
	result := model.DB.Create(&group)
	if result.Error != nil {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "New Group Name Must Be Unique",
		}, nil
	}
	in.Data["groupid"] = group.ID
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
