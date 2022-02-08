package socket

import (
	"errors"
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm"
)

func GroupDelete(in *system.WebSocketRequest) (*system.WebSocketReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.GroupDelete(in *system.WebSocketRequest) (*system.WebSocketReturn, error)]")
	var group model.Group
	if err := model.DB.First(&group, in.Data["groupid"]).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusNotFound,
				ErrorMessage: "Group Not Found",
			}, nil
		}
		return &system.WebSocketReturn{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "DB Group ERROR: " + err.Error(),
		}, nil
	}
	if err := model.DB.Delete(&group).Error; err != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "DB Group Delete ERROR: " + err.Error(),
		}, nil
	}
	return &system.WebSocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
