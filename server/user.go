package server

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	pb "github.com/Tackem-org/Proto/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
	"gorm.io/gorm/clause"
)

func (u *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[server.(u *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error)]")
	var user model.User
	result := model.DB.Where(&model.User{Username: in.Username, Password: password.Hash(in.Password)}).Find(&user)
	if result.RowsAffected == 1 {
		session := newSession(user.ID, in.GetIpAddress(), time.Duration(in.GetExpiryTime()))
		return &pb.LoginResponse{
			Success:      true,
			SessionToken: session,
		}, nil
	}

	return &pb.LoginResponse{
		Success:      false,
		ErrorMessage: "User Not Found or incorrect Password",
	}, nil
}

func (u *UserServer) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[server.(u *UserServer) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {]")
	for index, s := range Sessions {
		if s.SessionToken == in.SessionToken && s.IPAddress == in.IpAddress {
			Sessions = append(Sessions[:index], Sessions[index+1:]...)
			return &pb.LogoutResponse{
				Success: true,
			}, nil
		}
	}
	return &pb.LogoutResponse{
		Success:      false,
		ErrorMessage: "Session Not Found",
	}, nil
}

func (u *UserServer) GetUserData(ctx context.Context, in *pb.GetUserDataRequest) (*pb.UserDataResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[server.(u *UserServer) GetUserData(context.Context, *GetUserDataRequest) (*GetBaseDataResponse, error)]")
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
