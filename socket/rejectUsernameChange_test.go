package socket_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/socket"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestRejectUsernameChange(t *testing.T) {
	pflag.Set("config", "")
	model.Setup("testRejectUsernameChange.db")
	defer os.Remove("testRejectUsernameChange.db")
	model.DB.Create(&model.UsernameRequest{RequestUserID: 2, Name: "bob"})

	sr1, err1 := socket.RejectUsernameChange(&structs.SocketRequest{
		Command: "",
		User:    &structs.UserData{},
		Data:    map[string]interface{}{},
	})
	assert.IsType(t, &structs.SocketReturn{}, sr1)
	assert.Nil(t, err1)
	assert.Equal(t, http.StatusBadRequest, int(sr1.StatusCode))
	assert.Equal(t, "userid not valid", sr1.ErrorMessage)

	sr2, err2 := socket.RejectUsernameChange(&structs.SocketRequest{
		Command: "",
		User:    &structs.UserData{},
		Data: map[string]interface{}{
			"userid": 3,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, sr2)
	assert.Nil(t, err2)
	assert.Equal(t, http.StatusBadRequest, int(sr2.StatusCode))
	assert.Equal(t, "userid not found", sr2.ErrorMessage)

	sr3, err3 := socket.RejectUsernameChange(&structs.SocketRequest{
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

	sr4, err4 := socket.RejectUsernameChange(&structs.SocketRequest{
		Command: "",
		User:    &structs.UserData{},
		Data: map[string]interface{}{
			"userid": 2,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, sr4)
	assert.Nil(t, err4)
	assert.Equal(t, http.StatusOK, int(sr4.StatusCode))
	assert.Empty(t, sr4.ErrorMessage)
	os.Remove("Salt.dat")
	os.Remove("adminpassword")
}
