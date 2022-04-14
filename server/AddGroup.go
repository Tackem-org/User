package server

import (
	"context"

	pb "github.com/Tackem-org/Global/pb/user"
	"github.com/Tackem-org/User/model"
)

func (u *UserServer) AddGroups(ctx context.Context, in *pb.AddGroupsRequest) (*pb.AddGroupsResponse, error) {
	model.AddGroups(in.Groups...)

	return &pb.AddGroupsResponse{
		Success: true,
	}, nil
}
