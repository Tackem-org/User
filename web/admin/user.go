package admin

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"net/http"
	"regexp"
	"strings"

	"github.com/Tackem-org/Global/config"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
	"golang.org/x/image/draw"
	"gorm.io/gorm/clause"
)

func AdminUserIDPage(in *system.WebRequest) (*system.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminUserIDPage(in *system.WebRequest) (*system.WebReturn, error)]")

	var user model.User
	model.DB.Preload(clause.Associations).First(&user, in.PathVariables["userid"])

	var allPermissions []model.Permission
	model.DB.Find(&allPermissions)
	var allPermissionsList []sPermissions
	for _, permission := range allPermissions {
		p := sPermissions{
			ID:    permission.ID,
			Name:  permission.Name,
			Title: strings.ReplaceAll(permission.Name, "_", " "),
		}
		for _, v := range user.Permissions {
			if v.Name == permission.Name {
				p.Active = true
				break
			}
		}
		allPermissionsList = append(allPermissionsList, p)
	}

	var allGroups []model.Group
	model.DB.Find(&allGroups)
	var allGroupsList []sGroups
	for _, group := range allGroups {
		g := sGroups{
			ID:    group.ID,
			Name:  group.Name,
			Title: strings.ReplaceAll(group.Name, "_", " "),
		}
		for _, v := range user.Groups {
			if v.Name == group.Name {
				g.Active = true
				break
			}
		}
		allGroupsList = append(allGroupsList, g)
	}

	return &system.WebReturn{
		StatusCode:     http.StatusOK,
		FilePath:       "admin/user",
		CustomPageName: "admin-user-edit",
		PageData: map[string]interface{}{
			"User":        user,
			"Permissions": allPermissionsList,
			"Groups":      allGroupsList,
		},
	}, nil
}

func AdminEditUserWebSocket(in *system.WebSocketRequest) (*system.WebSocketReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[web.AdminEditUserWebSocket(in *system.WebSocketRequest) (*system.WebSocketReturn, error)]")
	d := in.Data
	userID := d["userid"]
	var user model.User
	result := model.DB.Preload(clause.Associations).Find(&user, userID)
	if result.Error != nil {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "user not found",
		}, nil
	}

	command, ok := d["command"].(string)
	if !ok {
		return &system.WebSocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "COMMAND NOT FOUND",
		}, nil
	}
	switch command {
	case "changeusername":
		val, ok := d["username"].(string) //TODO finish this down
		if !ok || val == "" || len(val) <= 4 || !regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString(val) {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "username not valid",
			}, nil
		}
		// user.Username = val
		// result := model.DB.Save(&user)
		result := model.DB.Model(&user).Update("Username", val)
		if result.Error != nil {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "username already exists " + result.Error.Error(),
			}, nil
		}
	case "changepassword":
		val, ok := d["password"].(string)
		if !ok {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "password not valid",
			}, nil
		}
		minPassLength, _ := config.GetUint("user.password.minimum")
		if uint(len(val)) <= minPassLength {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "password too short",
			}, nil
		}
		// user.Password = password.Hash(val)
		// result := model.DB.Save(&user)
		result := model.DB.Model(&user).Update("Password", password.Hash(val))
		if result.Error != nil {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "password Error " + result.Error.Error(),
			}, nil
		}
	case "changedisabled":
		val, ok := d["checked"].(bool)
		if !ok {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "changing disabled failed",
			}, nil
		}
		// user.Disabled = val
		// result := model.DB.Save(&user)
		result := model.DB.Model(&user).Update("Disabled", val)
		if result.Error != nil {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "disabled error " + result.Error.Error(),
			}, nil
		}
	case "changeisadmin":
		val, ok := d["checked"].(bool)
		if !ok {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "changing disabled failed",
			}, nil
		}
		// user.IsAdmin = val
		// result := model.DB.Save(&user)
		result := model.DB.Model(&user).Update("IsAdmin", val)
		if result.Error != nil {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "isadmin error " + result.Error.Error(),
			}, nil
		}
	case "uploadiconbase64":
		val, ok := d["icon"].(string)
		if !ok || val == "" {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "uploading icon failed",
			}, nil
		}
		b64data := val[strings.IndexByte(val, ',')+1:]
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64data))
		src, _, err := image.Decode(reader)
		if err != nil {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "uploading icon failed cannot decode error",
			}, nil
		}
		dst := image.NewRGBA(image.Rect(0, 0, 64, 64))
		draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
		buf := new(bytes.Buffer)
		if err := png.Encode(buf, dst); err != nil {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "uploading icon failed cannot encode to png",
			}, nil
		}
		imgBase64Str := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buf.Bytes()))
		d["icon"] = imgBase64Str
		// user.Icon = imgBase64Str
		// result := model.DB.Save(&user)
		result := model.DB.Model(&user).Update("Icon", imgBase64Str)
		if result.Error != nil {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "uploading icon error " + result.Error.Error(),
			}, nil
		}
	case "clearicon":
		// user.Icon = ""
		// result := model.DB.Save(&user)
		result := model.DB.Model(&user).Update("Icon", "")
		if result.Error != nil {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "clearing icon error " + result.Error.Error(),
			}, nil
		}
	case "deleteuser":
		result := model.DB.Delete(&user)
		if result.Error != nil {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: "delete error " + result.Error.Error(),
			}, nil
		}
	case "changegroup":
		var group model.Group
		result := model.DB.First(&group, d["group"])
		if result.Error != nil {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusNotFound,
				ErrorMessage: "group not found",
			}, nil
		}
		if d["checked"] == true {
			model.DB.Model(&user).Association("Groups").Append(&group)
		} else {
			model.DB.Model(&user).Association("Groups").Delete(&group)
		}
	case "changepermission":
		var permission model.Permission
		result := model.DB.First(&permission, d["permission"])
		if result.Error != nil {
			return &system.WebSocketReturn{
				StatusCode:   http.StatusNotFound,
				ErrorMessage: "permission not found",
			}, nil
		}
		if d["checked"] == true {
			model.DB.Model(&user).Association("Permissions").Append(&permission)
		} else {
			model.DB.Model(&user).Association("Permissions").Delete(&permission)
		}
	default:
		return &system.WebSocketReturn{
			StatusCode:   http.StatusNotImplemented,
			ErrorMessage: "command not found: " + command,
		}, nil
	}

	d["updatedat"] = user.UpdatedAt
	return &system.WebSocketReturn{
		StatusCode: http.StatusOK,
		Data:       d,
	}, nil
}
