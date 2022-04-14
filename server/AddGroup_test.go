package server_test

import (
	"context"
	"testing"

	pb "github.com/Tackem-org/Global/pb/user"
	"github.com/Tackem-org/User/server"
	"github.com/stretchr/testify/assert"
)

func TestUserServerAddGroups(t *testing.T) {
	u := server.UserServer{}
	response, err := u.AddGroups(context.Background(), &pb.AddGroupsRequest{Groups: []string{}})
	assert.True(t, response.Success)
	assert.Nil(t, err)
}
