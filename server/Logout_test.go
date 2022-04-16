package server_test

import (
	"context"
	"os"
	"testing"
	"time"

	pb "github.com/Tackem-org/Global/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/server"
	"github.com/stretchr/testify/assert"
)

func TestUserServer_Logout(t *testing.T) {
	u := server.UserServer{}
	assert.NotPanics(t, func() { model.Setup("testServerLogout.db") })
	defer os.Remove("testServerLogout.db")
	server.Sessions = []server.Session{
		{
			UserID:       2,
			SessionToken: "pass",
			IPAddress:    "127.0.0.1",
			ExpireTime:   time.Now().Add(time.Second),
		},
	}

	response1, err1 := u.Logout(context.Background(), &pb.LogoutRequest{
		SessionToken: "fail",
		IpAddress:    "127.0.0.1",
	})
	assert.IsType(t, &pb.LogoutResponse{}, response1)
	assert.False(t, response1.Success)
	assert.Nil(t, err1)
	response2, err2 := u.Logout(context.Background(), &pb.LogoutRequest{
		SessionToken: "pass",
		IpAddress:    "127.0.0.1",
	})
	assert.IsType(t, &pb.LogoutResponse{}, response2)
	assert.True(t, response2.Success)
	assert.Nil(t, err2)
}
