package admin_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/web/admin"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

func TestAdminUserIDPage(t *testing.T) {
	pflag.Set("config", "")
	model.Setup("testAdminUserIDPage.db")
	defer os.Remove("testAdminUserIDPage.db")

	var user model.User
	model.DB.Preload(clause.Associations).Where(&model.User{ID: uint64(1)}).Find(&user)
	var permission model.Permission
	model.DB.Where(&model.Permission{ID: uint64(1)}).First(&permission)
	model.DB.Model(&user).Association("Permissions").Append(&permission)

	var group model.Group
	model.DB.Where(&model.Group{ID: uint64(1)}).First(&group)
	model.DB.Model(&user).Association("Groups").Append(&group)

	r1, err1 := admin.AdminUserIDPage(&structs.WebRequest{
		PathVariables: map[string]interface{}{},
	})
	assert.IsType(t, &structs.WebReturn{}, r1)
	assert.Nil(t, err1)
	assert.Equal(t, http.StatusNotAcceptable, int(r1.StatusCode))
	assert.Equal(t, "userid missing", r1.ErrorMessage)

	r2, err2 := admin.AdminUserIDPage(&structs.WebRequest{
		PathVariables: map[string]interface{}{
			"userid": "fail",
		},
	})
	assert.IsType(t, &structs.WebReturn{}, r2)
	assert.Nil(t, err2)
	assert.Equal(t, http.StatusNotAcceptable, int(r2.StatusCode))
	assert.Equal(t, "userid not an int", r2.ErrorMessage)

	r3, err3 := admin.AdminUserIDPage(&structs.WebRequest{
		PathVariables: map[string]interface{}{
			"userid": 30,
		},
	})
	assert.IsType(t, &structs.WebReturn{}, r3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusNotFound, int(r3.StatusCode))
	assert.Equal(t, "user not found", r3.ErrorMessage)

	r4, err4 := admin.AdminUserIDPage(&structs.WebRequest{
		PathVariables: map[string]interface{}{
			"userid": 1,
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r4)
	assert.Nil(t, err4)
	assert.Equal(t, http.StatusOK, int(r4.StatusCode))
	assert.Equal(t, "", r4.ErrorMessage)
	os.Remove("Salt.dat")
	os.Remove("adminpassword")
}
