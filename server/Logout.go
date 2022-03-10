package server

import (
	"context"

	pb "github.com/Tackem-org/Proto/pb/user"
)

func (u *UserServer) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
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
