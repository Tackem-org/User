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
			ErrorMessage: "name missing",
		}, nil
	}

	name := tmpName.(string)
	if name == "" {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "new group name cannot be blank",
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
			ErrorMessage: "new group name must be unique",
		}, nil
	}
	model.DB.Create(&group)
	in.Data["groupid"] = group.ID
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
