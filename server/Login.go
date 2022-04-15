package server

import (
	"context"
	"time"

	pb "github.com/Tackem-org/Global/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
	"github.com/google/uuid"
)

func (u *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user model.User
	result := model.DB.Where(&model.User{Username: in.Username, Password: password.Hash(in.Password)}).Find(&user)
	if result.RowsAffected == 1 {
		session := NewSession(user.ID, in.GetIpAddress(), time.Duration(in.GetExpiryTime()))
		return &pb.LoginResponse{
			Success:      true,
			SessionToken: session,
		}, nil
	}

	return &pb.LoginResponse{
		Success:      false,
		ErrorMessage: "user not found or incorrect password",
	}, nil
}

func NewSession(userID uint64, ipAddress string, expiryTime time.Duration) string {
	new := Session{
		UserID:       userID,
		SessionToken: uuid.New().String(),
		IPAddress:    ipAddress,
		ExpireTime:   time.Now().Add(expiryTime),
	}
	Sessions = append(Sessions, new)
	return new.SessionToken
}
