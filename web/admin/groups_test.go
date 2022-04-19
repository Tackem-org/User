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

func TestAdminGroupsPage(t *testing.T) {
	pflag.Set("config", "")
	model.Setup("testAdminGroupsPage.db")
	defer os.Remove("testAdminGroupsPage.db")

	r1, err1 := admin.AdminGroupsPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{},
	})

	assert.IsType(t, &structs.WebReturn{}, r1)
	assert.Nil(t, err1)
	os.Remove("Salt.dat")
	os.Remove("adminpassword")
}
