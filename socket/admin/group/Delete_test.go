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

func TestDelete(t *testing.T) {
	model.Setup("testDelete.db")
	defer os.Remove("testDelete.db")
	gt := &model.Group{Name: "test1"}
	model.DB.Create(gt)

	r1, err1 := group.Delete(&structs.SocketRequest{
		Data: map[string]interface{}{},
	})
	assert.IsType(t, &structs.SocketReturn{}, r1)
	assert.Nil(t, err1)
	assert.Equal(t, http.StatusNotAcceptable, int(r1.StatusCode))
	assert.Equal(t, "groupid missing", r1.ErrorMessage)

	r2, err2 := group.Delete(&structs.SocketRequest{
		Data: map[string]interface{}{
			"groupid": 0,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r2)
	assert.Nil(t, err2)
	assert.Equal(t, http.StatusNotAcceptable, int(r2.StatusCode))
	assert.Equal(t, "groupid cannot be zero", r2.ErrorMessage)

	r3, err3 := group.Delete(&structs.SocketRequest{
		Data: map[string]interface{}{
			"groupid": 999,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusNotFound, int(r3.StatusCode))
	assert.Equal(t, "group not found", r3.ErrorMessage)

	r4, err4 := group.Delete(&structs.SocketRequest{
		Data: map[string]interface{}{
			"groupid": 2,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r4)
	assert.Nil(t, err4)
	assert.Equal(t, http.StatusOK, int(r4.StatusCode))
	assert.Empty(t, r4.ErrorMessage)
}
