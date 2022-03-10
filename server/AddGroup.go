package server

import (
	"context"

	pb "github.com/Tackem-org/Proto/pb/user"
	"github.com/Tackem-org/User/model"
)

func (u *UserServer) AddGroups(ctx context.Context, in *pb.AddGroupsRequest) (*pb.AddGroupsResponse, error) {
	for _, group := range in.Groups {
		model.AddGroup(group)
	}
	return &pb.AddGroupsResponse{
		Success: true,
	}, nil
}
