package editUser

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"net/http"
	"strings"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"golang.org/x/image/draw"
	"gorm.io/gorm/clause"
)

func UploadIconBase64(in *structs.SocketRequest) (*structs.SocketReturn, error) {
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

	tmpIcon, foundIcon := in.Data["icon"]
	if !foundIcon {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotAcceptable,
			ErrorMessage: "icon missing",
		}, nil
	}
	val, ok := tmpIcon.(string)
	if !ok {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "icon not a string",
		}, nil
	}

	b64data := val[strings.IndexByte(val, ',')+1:]
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64data))
	src, _, err := image.Decode(reader)

	if err != nil {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "icon not base64 encoded",
		}, nil
	}

	dst := image.NewRGBA(image.Rect(0, 0, 64, 64))
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	buf := new(bytes.Buffer)
	png.Encode(buf, dst)
	imgBase64Str := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buf.Bytes()))
	in.Data["icon"] = imgBase64Str
	model.DB.Model(&user).Update("Icon", imgBase64Str)
	in.Data["updatedat"] = user.UpdatedAt
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
