package web_test

import (
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/web"
	"github.com/stretchr/testify/assert"
)

func TestRootPage(t *testing.T) {
	r1, err1 := web.RootPage(&structs.WebRequest{})
	assert.IsType(t, &structs.WebReturn{}, r1)
	assert.Nil(t, err1)
}
