package tasks_test

import (
	"testing"
	"time"

	pbw "github.com/Tackem-org/Global/pb/web"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/tasks"
	"github.com/stretchr/testify/assert"
)

func TestUserNameChangeRequest(t *testing.T) {
	response := tasks.UserNameChangeRequest(&model.UsernameRequest{
		ID:            1,
		CreatedAt:     time.Now().Add(-time.Second),
		UpdatedAt:     time.Now().Add(-time.Second),
		RequestUserID: 2,
		Name:          "test",
	})
	assert.IsType(t, &pbw.TaskMessage{}, response)
}
