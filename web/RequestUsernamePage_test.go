package web_test

import (
	"os"
	"testing"

	pbw "github.com/Tackem-org/Global/pb/web"
	"github.com/Tackem-org/Global/structs"
	webClient "github.com/Tackem-org/Global/system/grpcSystem/clients/web"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/web"
	"github.com/stretchr/testify/assert"
)

type MockWebClient struct{}

func (mwc *MockWebClient) AddTask(request *pbw.TaskMessage) bool                { return true }
func (mwc *MockWebClient) RemoveTask(request *pbw.RemoveTaskRequest) bool       { return true }
func (mwc *MockWebClient) WebSocketSend(request *pbw.SendWebSocketRequest) bool { return true }

func TestRequestUsernamePage(t *testing.T) {
	webClient.I = &MockWebClient{}
	assert.NotPanics(t, func() { model.Setup("testRequestUsernamePage.db") })
	defer os.Remove("testRequestUsernamePage.db")

	var count int64
	model.DB.Model(&model.UsernameRequest{}).Count(&count)
	assert.Zero(t, count, count)

	r1, err1 := web.RequestUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2, Name: "user"},
		Post: map[string]interface{}{},
	})

	assert.IsType(t, &structs.WebReturn{}, r1)
	assert.Nil(t, err1)
	assert.False(t, r1.PageData["RequestMade"].(bool))
	assert.False(t, r1.PageData["Success"].(bool))
	assert.Equal(t, "", r1.PageData["Error"].(string))

	model.DB.Model(&model.UsernameRequest{}).Count(&count)
	assert.Zero(t, count, count)

	r2, err2 := web.RequestUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2, Name: "user"},
		Post: map[string]interface{}{
			"test": "test",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r2)
	assert.Nil(t, err2)
	assert.False(t, r2.PageData["RequestMade"].(bool))
	assert.False(t, r2.PageData["Success"].(bool))
	assert.Equal(t, "error cannot get post data", r2.PageData["Error"].(string))

	r3, err3 := web.RequestUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2, Name: "user"},
		Post: map[string]interface{}{
			"username": "test",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r3)
	assert.Nil(t, err3)
	assert.False(t, r3.PageData["RequestMade"].(bool))
	assert.False(t, r3.PageData["Success"].(bool))
	assert.Equal(t, "error cannot get post data", r3.PageData["Error"].(string))

	r4, err4 := web.RequestUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2, Name: "user"},
		Post: map[string]interface{}{
			"password": "test",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r4)
	assert.Nil(t, err4)
	assert.False(t, r4.PageData["RequestMade"].(bool))
	assert.False(t, r4.PageData["Success"].(bool))
	assert.Equal(t, "error cannot get post data", r4.PageData["Error"].(string))

	r5, err5 := web.RequestUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2, Name: "user"},
		Post: map[string]interface{}{
			"username": "",
			"password": "test",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r5)
	assert.Nil(t, err5)
	assert.False(t, r5.PageData["RequestMade"].(bool))
	assert.False(t, r5.PageData["Success"].(bool))
	assert.Equal(t, "new username cannot be blank", r5.PageData["Error"].(string))

	r6, err6 := web.RequestUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2, Name: "user"},
		Post: map[string]interface{}{
			"username": "user",
			"password": "",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r6)
	assert.Nil(t, err6)
	assert.False(t, r6.PageData["RequestMade"].(bool))
	assert.False(t, r6.PageData["Success"].(bool))
	assert.Equal(t, "username same as previous", r6.PageData["Error"].(string))

	r7, err7 := web.RequestUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2, Name: "user"},
		Post: map[string]interface{}{
			"username": "test",
			"password": "",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r7)
	assert.Nil(t, err7)
	assert.False(t, r7.PageData["RequestMade"].(bool))
	assert.False(t, r7.PageData["Success"].(bool))
	assert.Equal(t, "password cannot be blank", r7.PageData["Error"].(string))

	r8, err8 := web.RequestUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2, Name: "user"},
		Post: map[string]interface{}{
			"username": "admin",
			"password": "password",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r8)
	assert.Nil(t, err8)
	assert.False(t, r8.PageData["RequestMade"].(bool))
	assert.False(t, r8.PageData["Success"].(bool))
	assert.Equal(t, "user already exists", r8.PageData["Error"].(string))

	r9, err9 := web.RequestUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2, Name: "user"},
		Post: map[string]interface{}{
			"username": "test",
			"password": "password",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r9)
	assert.Nil(t, err9)
	assert.False(t, r9.PageData["RequestMade"].(bool))
	assert.False(t, r9.PageData["Success"].(bool))
	assert.Equal(t, "password doesn't match", r9.PageData["Error"].(string))

	r10, err10 := web.RequestUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2, Name: "user"},
		Post: map[string]interface{}{
			"username": "test",
			"password": "user",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r10)
	assert.Nil(t, err10)
	assert.False(t, r10.PageData["RequestMade"].(bool))
	assert.True(t, r10.PageData["Success"].(bool))
	assert.Equal(t, "", r10.PageData["Error"].(string))

	r11, err11 := web.RequestUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2, Name: "user"},
		Post: map[string]interface{}{},
	})

	assert.IsType(t, &structs.WebReturn{}, r11)
	assert.Nil(t, err11)
	assert.True(t, r11.PageData["RequestMade"].(bool))
	assert.False(t, r11.PageData["Success"].(bool))
	assert.Equal(t, "", r11.PageData["Error"].(string))
}
