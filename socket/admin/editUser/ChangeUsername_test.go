package editUser_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/password"
	"github.com/Tackem-org/User/socket/admin/editUser"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestChangeUsername(t *testing.T) {
	pflag.Set("config", "")
	model.Setup("testChangeUsername.db")
	defer os.Remove("testChangeUsername.db")
	model.DB.Create(&model.User{Username: "user", Password: password.Hash("user")})

	r1, err1 := editUser.ChangeUsername(&structs.SocketRequest{
		Data: map[string]interface{}{},
	})
	assert.IsType(t, &structs.SocketReturn{}, r1)
	assert.Nil(t, err1)
	assert.Equal(t, http.StatusNotAcceptable, int(r1.StatusCode))
	assert.Equal(t, "userid missing", r1.ErrorMessage)

	r2, err2 := editUser.ChangeUsername(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": "fail",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r2)
	assert.Nil(t, err2)
	assert.Equal(t, http.StatusNotAcceptable, int(r2.StatusCode))
	assert.Equal(t, "userid not an int", r2.ErrorMessage)

	r3, err3 := editUser.ChangeUsername(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 30,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusNotFound, int(r3.StatusCode))
	assert.Equal(t, "user not found", r3.ErrorMessage)

	r4, err4 := editUser.ChangeUsername(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 2,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r4)
	assert.Nil(t, err4)
	assert.Equal(t, http.StatusNotAcceptable, int(r4.StatusCode))
	assert.Equal(t, "username missing", r4.ErrorMessage)

	r5, err5 := editUser.ChangeUsername(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":   2,
			"username": false,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r5)
	assert.Nil(t, err5)
	assert.Equal(t, http.StatusBadRequest, int(r5.StatusCode))
	assert.Equal(t, "username not a string", r5.ErrorMessage)

	r6, err6 := editUser.ChangeUsername(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":   2,
			"username": "!!",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r6)
	assert.Nil(t, err6)
	assert.Equal(t, http.StatusBadRequest, int(r6.StatusCode))
	assert.Equal(t, "username not valid", r6.ErrorMessage)

	r7, err7 := editUser.ChangeUsername(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":   2,
			"username": "admin",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r7)
	assert.Nil(t, err7)
	assert.Equal(t, http.StatusBadRequest, int(r7.StatusCode))
	assert.Equal(t, "username already exists", r7.ErrorMessage)

	r8, err8 := editUser.ChangeUsername(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":   2,
			"username": "username123",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r8)
	assert.Nil(t, err8)
	assert.Equal(t, http.StatusOK, int(r8.StatusCode))
	os.Remove("Salt.dat")
	os.Remove("adminpassword")
}
