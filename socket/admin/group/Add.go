package group

import (
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
)

func Add(in *structs.SocketRequest) (*structs.SocketReturn, error) {
	tmpName, foundName := in.Data["name"]
	if !foundName {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "Name Missing",
		}, nil
	}

	name := tmpName.(string)
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
	var found model.Group
	model.DB.Where(&model.Group{Name: name}).First(&found)
	if found.ID > 0 {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "New Group Name Must Be Unique",
		}, nil
	}
	model.DB.Create(&group)
	in.Data["groupid"] = group.ID
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
