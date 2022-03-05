package group

import (
	"errors"
	"net/http"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm"
)

func Delete(in *structs.SocketRequest) (*structs.SocketReturn, error) {
	var group model.Group
	if err := model.DB.First(&group, in.Data["groupid"]).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &structs.SocketReturn{
				StatusCode:   http.StatusNotFound,
				ErrorMessage: "Group Not Found",
			}, nil
		}
		return &structs.SocketReturn{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "DB Group ERROR: " + err.Error(),
		}, nil
	}
	if err := model.DB.Delete(&group).Error; err != nil {
		return &structs.SocketReturn{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "DB Group Delete ERROR: " + err.Error(),
		}, nil
	}
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
