package server_test

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	pb "github.com/Tackem-org/Global/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/server"
	"github.com/stretchr/testify/assert"
)

func TestUserServerGetUserData(t *testing.T) {
	u := server.UserServer{}
	server.Sessions = []server.Session{
		{
			UserID:       3,
			SessionToken: "passuser1",
			IPAddress:    "127.0.0.1",
			ExpireTime:   time.Now().Add(time.Second),
		},
		{
			UserID:       4,
			SessionToken: "passuser2",
			IPAddress:    "127.0.0.2",
			ExpireTime:   time.Now().Add(time.Second),
		},
		{
			UserID:       5,
			SessionToken: "passuser3",
			IPAddress:    "127.0.0.3",
			ExpireTime:   time.Now().Add(time.Second),
		},
	}
	assert.NotPanics(t, func() { model.Setup("testUserServerGetUserData.db") })
	user1 := model.User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  "user1",
		Password:  "user1",
		Icon:      "data:test",
	}
	user2 := model.User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  "user2",
		Password:  "user2",
		Icon:      "icon.png",
	}
	user3 := model.User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  "user3",
		Password:  "user3",
		Icon:      "",
	}
	model.DB.Create(&user1)
	model.DB.Create(&user2)
	model.DB.Create(&user3)
	response1, err1 := u.GetUserData(context.Background(), &pb.GetUserDataRequest{
		SessionToken: "fail",
		IpAddress:    "127.0.0.1",
	})
	assert.IsType(t, &pb.UserDataResponse{}, response1)
	assert.False(t, response1.Success)
	assert.Nil(t, err1)

	response2, err2 := u.GetUserData(context.Background(), &pb.GetUserDataRequest{
		SessionToken: "passuser1",
		IpAddress:    "127.0.0.1",
	})
	assert.IsType(t, &pb.UserDataResponse{}, response2)
	assert.True(t, response2.Success)
	assert.Nil(t, err2)
	assert.True(t, strings.HasPrefix(response2.Icon, "data:"))

	response3, err3 := u.GetUserData(context.Background(), &pb.GetUserDataRequest{
		SessionToken: "passuser2",
		IpAddress:    "127.0.0.2",
	})
	assert.IsType(t, &pb.UserDataResponse{}, response3)
	assert.True(t, response3.Success)
	assert.Nil(t, err3)
	assert.True(t, strings.HasPrefix(response3.Icon, "user/static/img/icons/"))

	response4, err4 := u.GetUserData(context.Background(), &pb.GetUserDataRequest{
		SessionToken: "passuser3",
		IpAddress:    "127.0.0.3",
	})
	assert.IsType(t, &pb.UserDataResponse{}, response4)
	assert.True(t, response4.Success)
	assert.Nil(t, err4)
	assert.Empty(t, response4.Icon)

	os.Remove("testUserServerGetUserData.db")
}
