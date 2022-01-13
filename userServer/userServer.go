package userServer

import (
	"context"
	"time"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	pb "github.com/Tackem-org/Proto/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
	"github.com/google/uuid"
)

var (
	Sessions []session
)

type UserServer struct {
	pb.UnimplementedUserServer
}

type session struct {
	UserID       uint64
	SessionToken string
	IPAddress    string
	ExpireTime   time.Time
}

func (u *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[userServer.]")
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
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[userServer.]")
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
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[userServer.]")
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
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[userServer.]")
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

func (u *UserServer) IsAdmin(ctx context.Context, in *pb.IsAdminRequest) (*pb.IsAdminResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[userServer.]")
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

func newSession(userID uint64, ipAddress string, expiryTime time.Duration) string {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[userServer.newSession(userID uint64, iPAddress string, expiryTime time.Duration) string] {userID=%d, ipAddress=%s, expiryTime=%s}", userID, ipAddress, helpers.DurationToString(expiryTime))
	new := session{
		UserID:       userID,
		SessionToken: uuid.New().String(),
		IPAddress:    ipAddress,
		ExpireTime:   time.Now().Add(expiryTime),
	}
	Sessions = append(Sessions, new)
	return new.SessionToken
}
