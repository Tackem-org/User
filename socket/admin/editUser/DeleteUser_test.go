package editUser_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/socket/admin/editUser"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	assert.NotPanics(t, func() { model.Setup("testDeleteUser.db") })
	defer os.Remove("testDeleteUser.db")

	r1, err1 := editUser.DeleteUser(&structs.SocketRequest{
		Data: map[string]interface{}{},
	})
	assert.IsType(t, &structs.SocketReturn{}, r1)
	assert.Nil(t, err1)
	assert.Equal(t, http.StatusNotAcceptable, int(r1.StatusCode))
	assert.Equal(t, "userid missing", r1.ErrorMessage)

	r2, err2 := editUser.DeleteUser(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": "fail",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r2)
	assert.Nil(t, err2)
	assert.Equal(t, http.StatusNotAcceptable, int(r2.StatusCode))
	assert.Equal(t, "userid not an int", r2.ErrorMessage)

	r3, err3 := editUser.DeleteUser(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 30,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusNotFound, int(r3.StatusCode))
	assert.Equal(t, "user not found", r3.ErrorMessage)

	r4, err4 := editUser.DeleteUser(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 2,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r4)
	assert.Nil(t, err4)
	assert.Equal(t, http.StatusOK, int(r4.StatusCode))
}
