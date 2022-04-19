package editUser_test

import (
	"net/http"
	"os"
	"testing"

	pb "github.com/Tackem-org/Global/pb/config"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/socket/admin/editUser"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

type MockConfig struct{}

func (mc *MockConfig) Get(request *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	if request.Key == "user.password.minimum" {
		return &pb.GetConfigResponse{
			Success:      true,
			ErrorMessage: "",
			Key:          request.Key,
			Value:        "8",
			Type:         pb.ValueType_Uint,
		}, nil
	}
	return &pb.GetConfigResponse{
		Success:      false,
		ErrorMessage: "not found",
		Key:          request.Key,
	}, nil

}

func (mc *MockConfig) Set(request *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	return &pb.SetConfigResponse{
		Success:      true,
		ErrorMessage: "not found",
	}, nil
}

func TestChangePassword(t *testing.T) {
	pflag.Set("config", "")
	config.I = &MockConfig{}
	model.Setup("testChangePassword.db")
	defer os.Remove("testChangePassword.db")

	r1, err1 := editUser.ChangePassword(&structs.SocketRequest{
		Data: map[string]interface{}{},
	})
	assert.IsType(t, &structs.SocketReturn{}, r1)
	assert.Nil(t, err1)
	assert.Equal(t, http.StatusNotAcceptable, int(r1.StatusCode))
	assert.Equal(t, "userid missing", r1.ErrorMessage)

	r2, err2 := editUser.ChangePassword(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": "fail",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r2)
	assert.Nil(t, err2)
	assert.Equal(t, http.StatusNotAcceptable, int(r2.StatusCode))
	assert.Equal(t, "userid not an int", r2.ErrorMessage)

	r3, err3 := editUser.ChangePassword(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 30,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusNotFound, int(r3.StatusCode))
	assert.Equal(t, "user not found", r3.ErrorMessage)

	r4, err4 := editUser.ChangePassword(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 2,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r4)
	assert.Nil(t, err4)
	assert.Equal(t, http.StatusNotAcceptable, int(r4.StatusCode))
	assert.Equal(t, "password missing", r4.ErrorMessage)

	r5, err5 := editUser.ChangePassword(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":   2,
			"password": false,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r5)
	assert.Nil(t, err5)
	assert.Equal(t, http.StatusBadRequest, int(r5.StatusCode))
	assert.Equal(t, "password not a string", r5.ErrorMessage)

	r6, err6 := editUser.ChangePassword(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":   2,
			"password": "short",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r6)
	assert.Nil(t, err6)
	assert.Equal(t, http.StatusBadRequest, int(r6.StatusCode))
	assert.Equal(t, "password too short", r6.ErrorMessage)

	r7, err7 := editUser.ChangePassword(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":   2,
			"password": "password123",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r7)
	assert.Nil(t, err7)
	assert.Equal(t, http.StatusOK, int(r7.StatusCode))
	os.Remove("Salt.dat")
	os.Remove("adminpassword")
}
