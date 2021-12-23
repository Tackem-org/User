package userServer

import (
	"context"
	"time"

	pb "github.com/Tackem-org/Proto/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
	"github.com/google/uuid"
)

var (
	sessions []session
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

type session struct {
	UserID       uint64
	SessionToken string
	IPAddress    string
	ExpireTime   time.Time
}

func (u *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
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
	for index, s := range sessions {
		if s.SessionToken == in.SessionToken && s.IPAddress == in.IpAddress {
			sessions = append(sessions[:index], sessions[index+1:]...)
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
	for _, s := range sessions {
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
	for _, s := range sessions {
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
	for _, s := range sessions {
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

func newSession(userID uint64, iPAddress string, expiryTime time.Duration) string {
	new := session{
		UserID:       userID,
		SessionToken: uuid.New().String(),
		IPAddress:    iPAddress,
		ExpireTime:   time.Now().Add(expiryTime),
	}
	sessions = append(sessions, new)
	return new.SessionToken
}
