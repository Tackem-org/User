package web

import (
	"net/http"

	"github.com/Tackem-org/Global/config"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
)

func ChangePasswordPage(in *structs.WebRequest) (*structs.WebReturn, error) {
	minPassLength, _ := config.GetUint("user.password.minimum")
	success := false
	errorString := ""
	if len(in.Post) > 0 {
		op, ok1 := in.Post["op"].(string)
		np1, ok2 := in.Post["np1"].(string)
		np2, ok3 := in.Post["np2"].(string)

		if !ok1 || !ok2 || !ok3 {
			errorString = "error cannot get post data"
		} else if op == "" {
			errorString = "original password blank"
		} else if np1 == "" || np2 == "" {
			errorString = "new password cannot be blank"
		} else if uint(len(np1)) < minPassLength || uint(len(np2)) < minPassLength {
			errorString = "new password too short"
		} else if np1 != np2 {
			errorString = "new password don't match"
		} else {
			var user model.User
			model.DB.Where(&model.User{ID: in.User.ID, Password: password.Hash(op)}).First(&user)
			model.DB.Model(&user).Update("Password", password.Hash(np1))
			success = true
		}
	}
	return &structs.WebReturn{
		StatusCode: http.StatusOK,
		FilePath:   "changepassword",
		PageData: map[string]interface{}{
			"Success": success,
			"Error":   errorString,
		},
	}, nil
}
