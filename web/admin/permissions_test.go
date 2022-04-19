package admin_test

import (
	"os"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/web/admin"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestAdminPermissionsPage(t *testing.T) {
	pflag.Set("config", "")
	model.Setup("testAdminPermissionsPage.db")
	defer os.Remove("testAdminPermissionsPage.db")

	r1, err1 := admin.AdminPermissionsPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{},
	})

	assert.IsType(t, &structs.WebReturn{}, r1)
	assert.Nil(t, err1)
	os.Remove("Salt.dat")
	os.Remove("adminpassword")
}

func TestCheckActivePermissions(t *testing.T) {
	assert.False(t, admin.CheckActivePermissions(1, []model.Permission{}))
	assert.True(t, admin.CheckActivePermissions(1, []model.Permission{{ID: 1}}))
}
