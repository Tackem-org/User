package editUser

import (
	"net/http"
	"regexp"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

func ChangeUsername(in *structs.SocketRequest) (*structs.SocketReturn, error) {
	tmpUserID, foundUserID := in.Data["userid"]
	if !foundUserID {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "userid missing",
		}, nil
	}
	userID, okUserID := tmpUserID.(int)
	if !okUserID {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "userid not an int",
		}, nil
	}
	var user model.User
	model.DB.Preload(clause.Associations).Where(&model.User{ID: uint64(userID)}).Find(&user)
	if user.ID == 0 {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "user not found",
		}, nil
	}

	tmpUsername, foundUsername := in.Data["username"]
	if !foundUsername {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "username missing",
		}, nil
	}
	val, ok := tmpUsername.(string)
	if !ok {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "username not a string",
		}, nil
	}

	if !ok || val == "" || len(val) <= 4 || !regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString(val) {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "username not valid",
		}, nil
	}

	var count int64
	model.DB.Model(&model.User{}).Where(&model.User{Username: val}).Count(&count)
	if count > 0 {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "username already exists",
		}, nil
	}

	model.DB.Model(&user).Update("Username", val)
	in.Data["updatedat"] = user.UpdatedAt
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
