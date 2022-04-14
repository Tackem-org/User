package server

import (
	"context"

	pb "github.com/Tackem-org/Global/pb/user"
	"github.com/Tackem-org/User/model"
)

func (u *UserServer) AddPermissions(ctx context.Context, in *pb.AddPermissionsRequest) (*pb.AddPermissionsResponse, error) {
	model.AddPermissions(in.Permissions...)
	return &pb.AddPermissionsResponse{
		Success: true,
	}, nil
}
