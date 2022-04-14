package model_test

import (
	"os"
	"testing"

	"github.com/Tackem-org/User/model"
	"github.com/stretchr/testify/assert"
)

func TestUserAfterFind(t *testing.T) {
	assert.NotPanics(t, func() { model.Setup("testAfterFind.db") })
	user := model.User{}
	model.DB.First(&user, "1")
	assert.Empty(t, user.Password)
	model.DB.First(&user, "2")
	assert.Empty(t, user.Password)
	os.Remove("testAfterFind.db")
}

func TestUserAllPermissionStrings(t *testing.T) {
	assert.NotPanics(t, func() { model.Setup("testAllPermissionStrings.db") })
	user := model.User{}
	model.DB.First(&user, "2")
	group := model.Group{}
	model.DB.First(&group, "1")
	permission1 := model.Permission{}
	model.DB.First(&permission1, "1")
	permission2 := model.Permission{}
	model.DB.First(&permission2, "2")

	model.DB.Model(&group).Association("Permissions").Append(&permission1)
	model.DB.Model(&user).Association("Groups").Append(&group)
	model.DB.Model(&user).Association("Permissions").Append(&permission2)

	returnedPermissions := user.AllPermissionStrings()
	assert.Len(t, returnedPermissions, 2)
	os.Remove("testAllPermissionStrings.db")
}
