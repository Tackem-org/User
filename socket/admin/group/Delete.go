package group

import (
	"net/http"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
)

func Delete(in *structs.SocketRequest) (*structs.SocketReturn, error) {
	tmpGroupID, foundGroupID := in.Data["groupid"]
	if !foundGroupID {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "GroupID Missing",
		}, nil
	}

	groupID := tmpGroupID.(int)
	if groupID == 0 {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "GroupID Cannot Be Zero",
		}, nil
	}
	var group model.Group
	model.DB.Where(model.Group{ID: uint64(groupID)}).First(&group)
	if group.ID == 0 {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "Group Not Found",
		}, nil
	}
	model.DB.Delete(&group)
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
