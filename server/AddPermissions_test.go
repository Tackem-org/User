package server_test

import (
	"context"
	"testing"

	pb "github.com/Tackem-org/Global/pb/user"
	"github.com/Tackem-org/User/server"
	"github.com/stretchr/testify/assert"
)

func TestUserServerAddPermissions(t *testing.T) {
	u := server.UserServer{}
	response, err := u.AddPermissions(context.Background(), &pb.AddPermissionsRequest{Permissions: []string{}})
	assert.True(t, response.Success)
	assert.Nil(t, err)
}
