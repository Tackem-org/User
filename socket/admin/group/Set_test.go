package group_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/socket/admin/group"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	assert.NotPanics(t, func() { model.Setup("testSet.db") })

	r1, err1 := group.Set(&structs.SocketRequest{
		Data: map[string]interface{}{},
	})
	assert.IsType(t, &structs.SocketReturn{}, r1)
	assert.Nil(t, err1)
	assert.Equal(t, http.StatusNotAcceptable, int(r1.StatusCode))
	assert.Equal(t, "GroupID Missing", r1.ErrorMessage)

	r2, err2 := group.Set(&structs.SocketRequest{
		Data: map[string]interface{}{
			"groupid": 1,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r2)
	assert.Nil(t, err2)
	assert.Equal(t, http.StatusNotAcceptable, int(r2.StatusCode))
	assert.Equal(t, "PermissionID Missing", r2.ErrorMessage)

	r3, err3 := group.Set(&structs.SocketRequest{
		Data: map[string]interface{}{
			"groupid":      1,
			"permissionid": 1,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusNotAcceptable, int(r3.StatusCode))
	assert.Equal(t, "checked Missing", r3.ErrorMessage)

	r4, err4 := group.Set(&structs.SocketRequest{
		Data: map[string]interface{}{
			"groupid":      99,
			"permissionid": 99,
			"checked":      true,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r4)
	assert.Nil(t, err4)
	assert.Equal(t, http.StatusNotFound, int(r4.StatusCode))
	assert.Equal(t, "Group Not Found", r4.ErrorMessage)

	r5, err5 := group.Set(&structs.SocketRequest{
		Data: map[string]interface{}{
			"groupid":      1,
			"permissionid": 99,
			"checked":      true,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r5)
	assert.Nil(t, err5)
	assert.Equal(t, http.StatusNotFound, int(r5.StatusCode))
	assert.Equal(t, "Permission Not Found", r5.ErrorMessage)

	r6, err6 := group.Set(&structs.SocketRequest{
		Data: map[string]interface{}{
			"groupid":      1,
			"permissionid": 1,
			"checked":      true,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r6)
	assert.Nil(t, err6)
	assert.Equal(t, http.StatusOK, int(r6.StatusCode))

	r7, err7 := group.Set(&structs.SocketRequest{
		Data: map[string]interface{}{
			"groupid":      1,
			"permissionid": 1,
			"checked":      false,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r7)
	assert.Nil(t, err7)
	assert.Equal(t, http.StatusOK, int(r7.StatusCode))

	defer os.Remove("testSet.db")
}
