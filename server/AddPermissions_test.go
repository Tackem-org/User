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

func TestUserServerAddPermissions(t *testing.T) {
	u := server.UserServer{}
	model.Setup("testServerAddPermissions.db")
	defer os.Remove("testServerAddPermissions.db")
	response, err := u.AddPermissions(context.Background(), &pb.AddPermissionsRequest{Permissions: []string{}})
	assert.True(t, response.Success)
	assert.Nil(t, err)
}
