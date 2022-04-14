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

func TestAdd(t *testing.T) {
	assert.NotPanics(t, func() { model.Setup("testAdd.db") })
	model.DB.Create(&model.Group{Name: "existing"})

	r1, err1 := group.Add(&structs.SocketRequest{
		Data: map[string]interface{}{},
	})
	assert.IsType(t, &structs.SocketReturn{}, r1)
	assert.Nil(t, err1)
	assert.Equal(t, http.StatusNotAcceptable, int(r1.StatusCode))
	assert.Equal(t, "Name Missing", r1.ErrorMessage)

	r2, err2 := group.Add(&structs.SocketRequest{
		Data: map[string]interface{}{
			"name": "",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r2)
	assert.Nil(t, err2)
	assert.Equal(t, http.StatusNotAcceptable, int(r2.StatusCode))
	assert.Equal(t, "New Group Name Cannot Be Blank", r2.ErrorMessage)

	r3, err3 := group.Add(&structs.SocketRequest{
		Data: map[string]interface{}{
			"name": "existing",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusNotAcceptable, int(r3.StatusCode))
	assert.Equal(t, "New Group Name Must Be Unique", r3.ErrorMessage)

	r4, err4 := group.Add(&structs.SocketRequest{
		Data: map[string]interface{}{
			"name": "new",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r4)
	assert.Nil(t, err4)
	assert.Equal(t, http.StatusOK, int(r4.StatusCode))
	assert.Empty(t, r4.ErrorMessage)
	os.Remove("testAdd.db")
}
