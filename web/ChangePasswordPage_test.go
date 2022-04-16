package web_test

import (
	"os"
	"testing"

	pb "github.com/Tackem-org/Global/pb/config"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/config"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/web"
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

func TestChangePasswordPage(t *testing.T) {
	config.I = &MockConfig{}
	model.Setup("testChangePasswordPage.db")
	defer os.Remove("testChangePasswordPage.db")

	r1, err1 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{},
	})

	assert.IsType(t, &structs.WebReturn{}, r1)
	assert.Nil(t, err1)
	assert.False(t, r1.PageData["Success"].(bool))
	assert.Equal(t, "", r1.PageData["Error"].(string))

	r2, err2 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"test": "test",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r2)
	assert.Nil(t, err2)
	assert.False(t, r2.PageData["Success"].(bool))
	assert.Equal(t, "error cannot get post data", r2.PageData["Error"].(string))

	r3, err3 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"op": "",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r3)
	assert.Nil(t, err3)
	assert.False(t, r3.PageData["Success"].(bool))
	assert.Equal(t, "error cannot get post data", r3.PageData["Error"].(string))

	r4, err4 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"op":  "",
			"np1": "",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r4)
	assert.Nil(t, err4)
	assert.False(t, r4.PageData["Success"].(bool))
	assert.Equal(t, "error cannot get post data", r4.PageData["Error"].(string))

	r5, err5 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"op":  "",
			"np1": "",
			"np2": "",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r5)
	assert.Nil(t, err5)
	assert.False(t, r5.PageData["Success"].(bool))
	assert.Equal(t, "original password blank", r5.PageData["Error"].(string))

	r6, err6 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"op":  "user",
			"np1": "",
			"np2": "",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r6)
	assert.Nil(t, err6)
	assert.False(t, r6.PageData["Success"].(bool))
	assert.Equal(t, "new password cannot be blank", r6.PageData["Error"].(string))

	r7, err7 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"op":  "user",
			"np1": "notblank",
			"np2": "",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r7)
	assert.Nil(t, err7)
	assert.False(t, r7.PageData["Success"].(bool))
	assert.Equal(t, "new password cannot be blank", r7.PageData["Error"].(string))

	r8, err8 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"op":  "user",
			"np1": "",
			"np2": "notblank",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r8)
	assert.Nil(t, err8)
	assert.False(t, r8.PageData["Success"].(bool))
	assert.Equal(t, "new password cannot be blank", r8.PageData["Error"].(string))

	r9, err9 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"op":  "user",
			"np1": "password123",
			"np2": "short",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r9)
	assert.Nil(t, err9)
	assert.False(t, r9.PageData["Success"].(bool))
	assert.Equal(t, "new password too short", r9.PageData["Error"].(string))

	r10, err10 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"op":  "user",
			"np1": "short",
			"np2": "password123",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r10)
	assert.Nil(t, err10)
	assert.False(t, r10.PageData["Success"].(bool))
	assert.Equal(t, "new password too short", r10.PageData["Error"].(string))

	r11, err11 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"op":  "user",
			"np1": "password321",
			"np2": "password123",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r11)
	assert.Nil(t, err11)
	assert.False(t, r11.PageData["Success"].(bool))
	assert.Equal(t, "new password don't match", r11.PageData["Error"].(string))

	r12, err12 := web.ChangePasswordPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"op":  "user",
			"np1": "password123",
			"np2": "password123",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r12)
	assert.Nil(t, err12)
	assert.True(t, r12.PageData["Success"].(bool))
	assert.Equal(t, "", r12.PageData["Error"].(string))
}
