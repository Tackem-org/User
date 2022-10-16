package model_test

import (
	"os"
	"testing"

	"github.com/Tackem-org/User/model"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	pflag.Set("config", "")
	model.Setup("testSetup.db")
	defer os.Remove("testSetup.db")
	var count1 int64
	model.DB.Model(&model.User{}).Count(&count1)
	assert.Equal(t, int64(1), count1)
	model.DB.Model(&model.Permission{}).Count(&count1)
	assert.Equal(t, int64(6), count1)
	model.DB.Model(&model.Group{}).Count(&count1)
	assert.Equal(t, int64(3), count1)

	model.Setup("testSetup.db")
	var count2 int64
	model.DB.Model(&model.User{}).Count(&count2)
	assert.Equal(t, int64(1), count2)
	model.DB.Model(&model.Permission{}).Count(&count2)
	assert.Equal(t, int64(6), count2)
	model.DB.Model(&model.Group{}).Count(&count2)
	assert.Equal(t, int64(3), count2)
	os.Remove("Salt.dat")
	os.Remove("adminpassword")
}
