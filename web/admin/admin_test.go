package admin_test

import (
	"os"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/web/admin"
	"github.com/stretchr/testify/assert"
)

func TestAdminRootPage(t *testing.T) {
	model.Setup("testAdminRootPage.db")
	defer os.Remove("testAdminRootPage.db")

	r1, err1 := admin.AdminRootPage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{},
	})

	assert.IsType(t, &structs.WebReturn{}, r1)
	assert.Nil(t, err1)
}
