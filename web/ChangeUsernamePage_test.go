package web_test

import (
	"os"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/web"
	"github.com/stretchr/testify/assert"
)

func TestChangeUsernamePage(t *testing.T) {
	assert.NotPanics(t, func() { model.Setup("testChangeUsernamePage.db") })
	defer os.Remove("testChangeUsernamePage.db")

	r1, err1 := web.ChangeUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{},
	})

	assert.IsType(t, &structs.WebReturn{}, r1)
	assert.Nil(t, err1)
	assert.False(t, r1.PageData["Success"].(bool))
	assert.Equal(t, "", r1.PageData["Error"].(string))

	r2, err2 := web.ChangeUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"test": "test",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r2)
	assert.Nil(t, err2)
	assert.False(t, r2.PageData["Success"].(bool))
	assert.Equal(t, "error cannot get post data", r2.PageData["Error"].(string))

	r3, err3 := web.ChangeUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"username": "test",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r3)
	assert.Nil(t, err3)
	assert.False(t, r3.PageData["Success"].(bool))
	assert.Equal(t, "error cannot get post data", r3.PageData["Error"].(string))

	r4, err4 := web.ChangeUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"password": "test",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r4)
	assert.Nil(t, err4)
	assert.False(t, r4.PageData["Success"].(bool))
	assert.Equal(t, "error cannot get post data", r4.PageData["Error"].(string))

	r5, err5 := web.ChangeUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"username": "",
			"password": "test",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r5)
	assert.Nil(t, err5)
	assert.False(t, r5.PageData["Success"].(bool))
	assert.Equal(t, "new username cannot be blank", r5.PageData["Error"].(string))

	r6, err6 := web.ChangeUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"username": "test",
			"password": "",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r6)
	assert.Nil(t, err6)
	assert.False(t, r6.PageData["Success"].(bool))
	assert.Equal(t, "password cannot be blank", r6.PageData["Error"].(string))

	r7, err7 := web.ChangeUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"username": "admin",
			"password": "password",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r7)
	assert.Nil(t, err7)
	assert.False(t, r7.PageData["Success"].(bool))
	assert.Equal(t, "user already exists", r7.PageData["Error"].(string))

	r8, err8 := web.ChangeUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"username": "test",
			"password": "wrong",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r8)
	assert.Nil(t, err8)
	assert.False(t, r8.PageData["Success"].(bool))
	assert.Equal(t, "password doesn't match", r8.PageData["Error"].(string))

	r9, err9 := web.ChangeUsernamePage(&structs.WebRequest{
		User: &structs.UserData{ID: 2},
		Post: map[string]interface{}{
			"username": "test",
			"password": "user",
		},
	})

	assert.IsType(t, &structs.WebReturn{}, r9)
	assert.Nil(t, err9)
	assert.True(t, r9.PageData["Success"].(bool))
	assert.Equal(t, "", r9.PageData["Error"].(string))
}
