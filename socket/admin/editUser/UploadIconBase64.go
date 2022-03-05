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
	userID := in.Data["userid"]
	var user model.User
	result := model.DB.Preload(clause.Associations).Find(&user, userID)
	if result.Error != nil {
		return &structs.SocketReturn{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "user not found",
		}, nil
	}

	val, ok := in.Data["icon"].(string)
	if !ok || val == "" {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "uploading icon failed",
		}, nil
	}
	b64data := val[strings.IndexByte(val, ',')+1:]
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64data))
	src, _, err := image.Decode(reader)
	if err != nil {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "uploading icon failed cannot decode error",
		}, nil
	}
	dst := image.NewRGBA(image.Rect(0, 0, 64, 64))
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, dst); err != nil {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "uploading icon failed cannot encode to png",
		}, nil
	}
	imgBase64Str := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buf.Bytes()))
	in.Data["icon"] = imgBase64Str
	result2 := model.DB.Model(&user).Update("Icon", imgBase64Str)
	if result2.Error != nil {
		return &structs.SocketReturn{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "uploading icon error " + result2.Error.Error(),
		}, nil
	}

	in.Data["updatedat"] = user.UpdatedAt
	return &structs.SocketReturn{
		StatusCode: http.StatusOK,
		Data:       in.Data,
	}, nil
}
