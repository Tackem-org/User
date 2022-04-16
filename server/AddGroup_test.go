package server_test

import (
	"context"
	"os"
	"testing"

	pb "github.com/Tackem-org/Global/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/server"
	"github.com/stretchr/testify/assert"
)

func TestUserServerAddGroups(t *testing.T) {
	u := server.UserServer{}
	assert.NotPanics(t, func() { model.Setup("testServerAddGroups.db") })
	defer os.Remove("testServerAddGroups.db")
	response, err := u.AddGroups(context.Background(), &pb.AddGroupsRequest{Groups: []string{}})
	assert.True(t, response.Success)
	assert.Nil(t, err)
}
