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

func TestChangePermission(t *testing.T) {
	assert.NotPanics(t, func() { model.Setup("testChangePermission.db") })
	defer os.Remove("testChangePermission.db")

	r1, err1 := editUser.ChangePermission(&structs.SocketRequest{
		Data: map[string]interface{}{},
	})
	assert.IsType(t, &structs.SocketReturn{}, r1)
	assert.Nil(t, err1)
	assert.Equal(t, http.StatusNotAcceptable, int(r1.StatusCode))
	assert.Equal(t, "userid missing", r1.ErrorMessage)

	r2, err2 := editUser.ChangePermission(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": "fail",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r2)
	assert.Nil(t, err2)
	assert.Equal(t, http.StatusNotAcceptable, int(r2.StatusCode))
	assert.Equal(t, "userid not an int", r2.ErrorMessage)

	r3, err3 := editUser.ChangePermission(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 30,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusNotFound, int(r3.StatusCode))
	assert.Equal(t, "user not found", r3.ErrorMessage)

	r4, err4 := editUser.ChangePermission(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 2,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r4)
	assert.Nil(t, err4)
	assert.Equal(t, http.StatusNotAcceptable, int(r4.StatusCode))
	assert.Equal(t, "permission missing", r4.ErrorMessage)

	r5, err5 := editUser.ChangePermission(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":     2,
			"permission": "fail",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r5)
	assert.Nil(t, err5)
	assert.Equal(t, http.StatusBadRequest, int(r5.StatusCode))
	assert.Equal(t, "permission not a int", r5.ErrorMessage)

	r6, err6 := editUser.ChangePermission(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":     2,
			"permission": 99,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r6)
	assert.Nil(t, err6)
	assert.Equal(t, http.StatusNotFound, int(r6.StatusCode))
	assert.Equal(t, "permission not found", r6.ErrorMessage)

	r7, err7 := editUser.ChangePermission(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":     2,
			"permission": 2,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r7)
	assert.Nil(t, err7)
	assert.Equal(t, http.StatusNotAcceptable, int(r7.StatusCode))
	assert.Equal(t, "checked missing", r7.ErrorMessage)

	r8, err8 := editUser.ChangePermission(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":     2,
			"permission": 2,
			"checked":    "fail",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r8)
	assert.Nil(t, err8)
	assert.Equal(t, http.StatusBadRequest, int(r8.StatusCode))
	assert.Equal(t, "checked not a bool", r8.ErrorMessage)

	r9, err9 := editUser.ChangePermission(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":     2,
			"permission": 2,
			"checked":    true,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r9)
	assert.Nil(t, err9)
	assert.Equal(t, http.StatusOK, int(r9.StatusCode))

	r10, err10 := editUser.ChangePermission(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid":     2,
			"permission": 2,
			"checked":    false,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r10)
	assert.Nil(t, err10)
	assert.Equal(t, http.StatusOK, int(r10.StatusCode))
}
