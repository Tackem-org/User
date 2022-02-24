package tasks

import (
	"fmt"

	"github.com/Tackem-org/Global/system"
	pbw "github.com/Tackem-org/Proto/pb/web"
	"github.com/Tackem-org/User/model"
)

type UsernameRequestCommand struct {
	Command   string `json:"command"`
	RequestID uint64 `json:"requestid"`
}

func UserNameChangeRequest(u *model.UsernameRequest) *pbw.TaskMessage {

	return &pbw.TaskMessage{
		Task:      "usernamechangerequest",
		BaseId:    system.GetBaseID(),
		TaskId:    u.ID,
		CreatedAt: u.CreatedAt.Format("2006-01-02T15:04:05"),
		Icon:      "user/static/img/userchange.png",
		Color:     "",
		Title:     fmt.Sprintf("Username Change Request '%s' -> '%s' ", u.RequestUser.Username, u.Name),
		Message:   fmt.Sprintf("User '%s' is requesting a change of there username to '%s'", u.RequestUser.Username, u.Name),
		Url:       "",
		Actions: []*pbw.TaskAction{
			{
				Title:   "Accept",
				Command: map[string]string{"command": "user.acceptusernamechange", "userid": fmt.Sprintf("%d", u.RequestUserID)},
			},
			{
				Title:   "Reject",
				Command: map[string]string{"command": "user.rejectusernamechange", "userid": fmt.Sprintf("%d", u.RequestUserID)},
			},
		},
		Who: &pbw.TaskMessage_Permission{
			Permission: "do_tasks",
		},
	}
}
