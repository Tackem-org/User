package server_test

import (
	"context"
	"os"
	"testing"

	pb "github.com/Tackem-org/Global/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/server"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestUserServerAddGroups(t *testing.T) {
	pflag.Set("config", "")
	u := server.UserServer{}
	model.Setup("testServerAddGroups.db")
	defer os.Remove("testServerAddGroups.db")
	response, err := u.AddGroups(context.Background(), &pb.AddGroupsRequest{Groups: []string{}})
	assert.True(t, response.Success)
	assert.Nil(t, err)
	os.Remove("Salt.dat")
	os.Remove("adminpassword")
}
