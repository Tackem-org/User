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

func (u *UserServer) Check(ctx context.Context, in *pb.CheckRequest) (*pb.CheckResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[server.(u *UserServer) Check(ctx context.Context, in *pb.CheckRequest) (*pb.CheckResponse, error)]")
	for _, s := range Sessions {
		if s.SessionToken == in.SessionToken && s.IPAddress == in.IpAddress {
			return &pb.CheckResponse{
				Success: true,
			}, nil
		}
	}
	return &pb.CheckResponse{
		Success:      false,
		ErrorMessage: "Session Not Found",
	}, nil
}

func (u *UserServer) GetUserID(ctx context.Context, in *pb.GetUserIDRequest) (*pb.GetUserIDResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[server.(u *UserServer) GetUserID(ctx context.Context, in *pb.GetUserIDRequest) (*pb.GetUserIDResponse, error)]")
	for _, s := range Sessions {
		if s.SessionToken == in.SessionToken && s.IPAddress == in.IpAddress {
			return &pb.GetUserIDResponse{
				Success: true,
				UserId:  s.UserID,
			}, nil
		}
	}
	return &pb.GetUserIDResponse{
		Success:      false,
		ErrorMessage: "Session Not Found",
	}, nil
}

func (u *UserServer) GetWebBaseData(ctx context.Context, in *pb.GetWebBaseDataRequest) (*pb.GetBaseDataResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[server.(u *UserServer) GetWebBaseData(context.Context, *GetWebBaseDataRequest) (*GetBaseDataResponse, error)]")
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
			return &pb.GetBaseDataResponse{
				Success:      true,
				ErrorMessage: "",
				UserId:       user.ID,
				Name:         user.Username,
				Initial:      strings.ToUpper(string(user.Username[0])),
				Icon:         icon,
				IsAdmin:      user.IsAdmin,
				Permissions:  user.AllPermissionStrings(),
			}, nil
		}
	}
	return &pb.GetBaseDataResponse{
		Success:      false,
		ErrorMessage: "Session Not Found",
	}, nil

}

func (u *UserServer) GetBaseData(ctx context.Context, in *pb.GetBaseDataRequest) (*pb.GetBaseDataResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[server.(u *UserServer) GetBaseData(context.Context, *GetBaseDataRequest) (*GetBaseDataResponse, error)]")

	var user model.User
	model.DB.Preload(clause.Associations).First(&user, in.UserId)
	var icon string
	if strings.HasPrefix(user.Icon, "data:") || strings.HasPrefix(user.Icon, "http") {
		icon = user.Icon
	} else if user.Icon != "" {
		icon = fmt.Sprintf("user/static/img/icons/%s", user.Icon)
	} else {
		icon = ""
	}
	return &pb.GetBaseDataResponse{
		Success:      true,
		ErrorMessage: "",
		UserId:       user.ID,
		Name:         user.Username,
		Initial:      strings.ToUpper(string(user.Username[0])),
		Icon:         icon,
		IsAdmin:      user.IsAdmin,
		Permissions:  user.AllPermissionStrings(),
	}, nil
}

func (u *UserServer) IsAdmin(ctx context.Context, in *pb.IsAdminRequest) (*pb.IsAdminResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[server.(u *UserServer) IsAdmin(ctx context.Context, in *pb.IsAdminRequest) (*pb.IsAdminResponse, error)]")
	for _, s := range Sessions {
		if s.SessionToken == in.SessionToken && s.IPAddress == in.IpAddress {
			var user model.User
			model.DB.First(&user, s.UserID)
			return &pb.IsAdminResponse{
				Success: true,
				IsAdmin: user.IsAdmin,
			}, nil
		}
	}
	return &pb.IsAdminResponse{
		Success:      false,
		ErrorMessage: "Session Not Found",
	}, nil
}
