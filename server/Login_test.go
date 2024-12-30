package server_test

import (
	"context"
	"os"
	"testing"
	"time"

	pb "github.com/Tackem-org/Global/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
	"github.com/Tackem-org/User/server"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestUserServerLogin(t *testing.T) {
	pflag.Set("config", "")
	u := server.UserServer{}
	server.Sessions = []server.Session{}
	model.Setup("testServerLogin.db")
	defer os.Remove("testServerLogin.db")

	model.DB.Create(&model.User{Username: "user", Password: password.Hash("user")})
	response1, err1 := u.Login(context.Background(), &pb.LoginRequest{
		Username:   "user",
		Password:   "fail",
		IpAddress:  "127.0.0.1",
		ExpiryTime: int64(time.Second),
	})
	assert.IsType(t, &pb.LoginResponse{}, response1)
	assert.False(t, response1.Success)
	assert.Nil(t, err1)
	response2, err2 := u.Login(context.Background(), &pb.LoginRequest{
		Username:   "user",
		Password:   "user",
		IpAddress:  "127.0.0.1",
		ExpiryTime: int64(time.Second),
	})
	assert.IsType(t, &pb.LoginResponse{}, response2)
	assert.True(t, response2.Success)
	assert.Nil(t, err2)
	os.Remove("Salt.dat")
	os.Remove("adminpassword")
}

func TestNewSession(t *testing.T) {
	server.Sessions = []server.Session{}
	assert.NotEmpty(t, server.NewSession(1, "127.0.0.1", time.Second))
	assert.Len(t, server.Sessions, 1)
	assert.IsType(t, server.Session{}, server.Sessions[0])
}
