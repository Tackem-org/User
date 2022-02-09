package group

import (
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
)

func Add(in *system.WebSocketRequest) (*system.WebSocketReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[socket.admin.group.GroupAdd(in *system.WebSocketRequest) (*system.WebSocketReturn, error)]")
	name := in.Data["name"].(string)
	if name == "" {
		return &system.WebSocketReturn{
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
		return &system.WebSocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "New Group Name Must Be Unique",
		}, nil
	}
	in.Data["groupid"] = group.ID
	return &system.WebSocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
