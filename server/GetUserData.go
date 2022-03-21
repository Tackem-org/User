package server

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/Tackem-org/Global/pb/user"
	"github.com/Tackem-org/User/model"
	"gorm.io/gorm/clause"
)

func (u *UserServer) GetUserData(ctx context.Context, in *pb.GetUserDataRequest) (*pb.UserDataResponse, error) {
	for _, s := range Sessions {
		if s.SessionToken == in.SessionToken && s.IPAddress == in.IpAddress {
			var user model.User
			model.DB.Preload(clause.Associations).First(&user, s.UserID)
			var icon string
			if strings.HasPrefix(user.Icon, "data:") || strings.HasPrefix(user.Icon, "http") {
				icon = user.Icon
			} else if user.Icon != "" {
				icon = fmt.Sprintf("user/static/img/icons/%s", user.Icon)
			} else {
				icon = ""
			}
			return &pb.UserDataResponse{
				Success:      true,
				ErrorMessage: "",
				UserId:       user.ID,
				Name:         user.Username,
				Icon:         icon,
				IsAdmin:      user.IsAdmin,
				Permissions:  user.AllPermissionStrings(),
			}, nil
		}
	}
	return &pb.UserDataResponse{
		Success:      false,
		ErrorMessage: "Session Not Found",
	}, nil

}
