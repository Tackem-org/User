package model_test

import (
	"os"
	"testing"

	"github.com/Tackem-org/User/model"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	assert.NotPanics(t, func() { model.Setup("testSetup.db") })
	var count1 int64
	model.DB.Model(&model.User{}).Count(&count1)
	assert.Equal(t, int64(2), count1)
	model.DB.Model(&model.Permission{}).Count(&count1)
	assert.Equal(t, int64(5), count1)
	model.DB.Model(&model.Group{}).Count(&count1)
	assert.Equal(t, int64(3), count1)

	assert.NotPanics(t, func() { model.Setup("testSetup.db") })
	var count2 int64
	model.DB.Model(&model.User{}).Count(&count2)
	assert.Equal(t, int64(2), count2)
	model.DB.Model(&model.Permission{}).Count(&count2)
	assert.Equal(t, int64(5), count2)
	model.DB.Model(&model.Group{}).Count(&count2)
	assert.Equal(t, int64(3), count2)
	os.Remove("testSetup.db")
}