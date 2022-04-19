package socket_test

import (
	"net/http"
	"os"
	"testing"

	pbw "github.com/Tackem-org/Global/pb/web"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/web"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/socket"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

type MockWebClient struct{}

func (wc *MockWebClient) AddTask(request *pbw.TaskMessage) bool                 { return true }
func (wc *MockWebClient) AddNotification(request *pbw.NotificationMessage) bool { return true }
func (wc *MockWebClient) RemoveTask(request *pbw.RemoveTaskRequest) bool        { return true }
func (wc *MockWebClient) WebSocketSend(request *pbw.SendWebSocketRequest) bool  { return true }

func TestAcceptUsernameChange(t *testing.T) {
	pflag.Set("config", "")
	web.I = &MockWebClient{}
	model.Setup("testAcceptUsernameChange.db")
	defer os.Remove("testAcceptUsernameChange.db")
	model.DB.Create(&model.User{Username: "test", Password: "test"})
	model.DB.Create(&model.UsernameRequest{RequestUserID: 2, Name: "bob"})
	model.DB.Create(&model.UsernameRequest{RequestUserID: 3, Name: "admin"})

	sr1, err1 := socket.AcceptUsernameChange(&structs.SocketRequest{
		Command: "",
		User:    &structs.UserData{},
		Data:    map[string]interface{}{},
	})
	assert.IsType(t, &structs.SocketReturn{}, sr1)
	assert.Nil(t, err1)
	assert.Equal(t, http.StatusBadRequest, int(sr1.StatusCode))
	assert.Equal(t, "userid not valid", sr1.ErrorMessage)

	sr2, err2 := socket.AcceptUsernameChange(&structs.SocketRequest{
		Command: "",
		User:    &structs.UserData{},
		Data: map[string]interface{}{
			"userid": 4,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, sr2)
	assert.Nil(t, err2)
	assert.Equal(t, http.StatusBadRequest, int(sr2.StatusCode))
	assert.Equal(t, "userid not found", sr2.ErrorMessage)

	sr3, err3 := socket.AcceptUsernameChange(&structs.SocketRequest{
		Command: "",
		User:    &structs.UserData{},
		Data: map[string]interface{}{
			"userid": 1,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, sr3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusBadRequest, int(sr3.StatusCode))
	assert.Equal(t, "username request not found", sr3.ErrorMessage)

	sr4, err4 := socket.AcceptUsernameChange(&structs.SocketRequest{
		Command: "",
		User:    &structs.UserData{},
		Data: map[string]interface{}{
			"userid": 3,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, sr4)
	assert.Nil(t, err4)
	assert.Equal(t, http.StatusBadRequest, int(sr4.StatusCode))
	assert.Equal(t, "username rename failed already exists", sr4.ErrorMessage)

	sr5, err5 := socket.AcceptUsernameChange(&structs.SocketRequest{
		Command: "",
		User:    &structs.UserData{},
		Data: map[string]interface{}{
			"userid": 2,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, sr5)
	assert.Nil(t, err5)
	assert.Equal(t, http.StatusOK, int(sr5.StatusCode))
	assert.Empty(t, sr5.ErrorMessage)
	os.Remove("Salt.dat")
	os.Remove("adminpassword")
}
