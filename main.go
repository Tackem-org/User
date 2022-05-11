package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/Tackem-org/Global/flags"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/Global/system/setupData"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/server"
	"github.com/Tackem-org/User/socket"
	"github.com/Tackem-org/User/socket/admin/editUser"
	"github.com/Tackem-org/User/socket/admin/group"
	"github.com/Tackem-org/User/static"
	"github.com/Tackem-org/User/tasks"
	"github.com/Tackem-org/User/web"
	"github.com/Tackem-org/User/web/admin"
	"google.golang.org/grpc"
	"gorm.io/gorm/clause"

	pbc "github.com/Tackem-org/Global/pb/config"
	pbr "github.com/Tackem-org/Global/pb/registration"
	pbu "github.com/Tackem-org/Global/pb/user"
	pbw "github.com/Tackem-org/Global/pb/web"
)

const (
	tempSaveFile     = "tackemusersessionsdata.tmp"
	masterConfigFile = "user.json"
	logFile          = "user.log"
	databaseFile     = "User.db"
)

var (
	Version    string = "v0.0.0-devel"
	Commit     string
	CommitDate string
	sd         = &setupData.SetupData{

		ServiceName: "user",
		ServiceType: "system",
		SingleRun:   false,
		StartActive: true,
		VerboseLog:  true,
		ConfigItems: []*pbr.ConfigItem{
			{
				Key:          "user.password.minimum",
				DefaultValue: "8",
				Type:         pbc.ValueType_Uint,
				Label:        "Password Minimum Length",
				HelpText:     "what is the minimum password length",
				InputType:    pbr.InputType_INumber,
				InputAttributes: &pbr.InputAttributes{
					Other: map[string]string{
						"min": "1",
						"max": "16",
					},
				},
			},
		},
		NavItems: []*pbr.NavItem{
			{LinkType: pbr.LinkType_User, Title: "Change Password", Icon: "user", Path: "/changepassword", Permission: "system_user_change_own_password"},
			{LinkType: pbr.LinkType_User, Title: "Change Username", Icon: "user", Path: "/changeusername", Permission: "system_user_change_own_username"},
			{LinkType: pbr.LinkType_User, Title: "Request New Username", Icon: "user", Path: "/requestusername", Permission: "system_user_request_change_of_username"},
			{LinkType: pbr.LinkType_Admin, Title: "Users", Icon: "users", Path: "/", SubLinks: []*pbr.NavItem{
				{LinkType: pbr.LinkType_Admin, Title: "Groups", Icon: "user-shield", Path: "/groups"},
				{LinkType: pbr.LinkType_Admin, Title: "Permissions", Icon: "key", Path: "/permissions"},
			},
			},
		},
		GRPCSystems: GRPCSystems,

		StaticFS: &static.FS,
		AdminPaths: []*setupData.AdminPathItem{
			{
				Path: "/",
				Call: admin.AdminRootPage,
			},
			{
				Path: "/edit/{{number:userid}}",
				Call: admin.AdminUserIDPage,
			},
			{
				Path: "/groups",
				Call: admin.AdminGroupsPage,
			},
			{
				Path: "/permissions",
				Call: admin.AdminPermissionsPage,
			},
		},
		Paths: []*setupData.PathItem{
			{
				Path:       "/",
				Permission: "",
				Call:       web.RootPage,
			},
			{
				Path:       "/changepassword",
				Permission: "system_user_change_own_password",
				Call:       web.ChangePasswordPage,
			},
			{
				Path:       "/changeusername",
				Permission: "system_user_change_own_username",
				Call:       web.ChangeUsernamePage,
			},
			{
				Path:       "/requestusername",
				Permission: "system_user_request_change_of_username",
				Call:       web.RequestUsernamePage,
			},
		},
		Sockets: []*setupData.SocketItem{
			{
				Command:           "admin.group.add",
				AdminOnly:         true,
				RequiredVariables: []string{"name"},
				Call:              group.Add,
			},
			{
				Command:           "admin.group.delete",
				AdminOnly:         true,
				RequiredVariables: []string{"groupid"},
				Call:              group.Delete,
			},
			{
				Command:           "admin.group.set",
				AdminOnly:         true,
				RequiredVariables: []string{"groupid", "permissionid", "checked"},
				Call:              group.Set,
			},
			{
				Command:           "admin.edituser.changeusername",
				AdminOnly:         true,
				RequiredVariables: []string{"userid", "username"},
				Call:              editUser.ChangeUsername,
			},
			{
				Command:           "admin.edituser.changepassword",
				AdminOnly:         true,
				RequiredVariables: []string{"userid", "password"},
				Call:              editUser.ChangePassword,
			},
			{
				Command:           "admin.edituser.changedisabled",
				AdminOnly:         true,
				RequiredVariables: []string{"userid", "checked"},
				Call:              editUser.ChangeDisabled,
			},
			{
				Command:           "admin.edituser.changeisadmin",
				AdminOnly:         true,
				RequiredVariables: []string{"userid", "checked"},
				Call:              editUser.ChangeIsAdmin,
			},
			{
				Command:           "admin.edituser.uploadiconbase64",
				AdminOnly:         true,
				RequiredVariables: []string{"userid", "icon"},
				Call:              editUser.UploadIconBase64,
			},
			{
				Command:           "admin.edituser.clearicon",
				AdminOnly:         true,
				RequiredVariables: []string{"userid"},
				Call:              editUser.ClearIcon,
			},
			{
				Command:           "admin.edituser.changegroup",
				AdminOnly:         true,
				RequiredVariables: []string{"userid", "group", "checked"},
				Call:              editUser.ChangeGroup,
			},
			{
				Command:           "admin.edituser.changepermission",
				AdminOnly:         true,
				RequiredVariables: []string{"userid", "permission", "checked"},
				Call:              editUser.ChangePermission,
			},
			{
				Command:           "acceptusernamechange",
				Permission:        "system_user_action_change_of_username",
				RequiredVariables: []string{"userid"},
				Call:              socket.AcceptUsernameChange,
			},
			{
				Command:           "rejectusernamechange",
				Permission:        "system_user_action_change_of_username",
				RequiredVariables: []string{"userid"},
				Call:              socket.RejectUsernameChange,
			},
		},
		TaskGrabber:         TaskGrabber,
		NotificationGrabber: NotificationGrabber,
		MainSetup:           MainSetup,
		MainShutdown:        MainShutdown,
	}
)

func main() {
	system.Run(sd, masterConfigFile, logFile, Version, Commit, CommitDate)
}

func TaskGrabber() []*pbw.TaskMessage {
	var rTasks []*pbw.TaskMessage
	var uChanges []model.UsernameRequest
	model.DB.Preload(clause.Associations).Find(&uChanges)
	for _, u := range uChanges {
		rTasks = append(rTasks, tasks.UserNameChangeRequest(&u))
	}
	return rTasks
}

func NotificationGrabber() []*pbw.NotificationMessage {
	var rNotifications []*pbw.NotificationMessage
	// TODO WHEN YOU HAVE NOTIFICATIONS MAKE THEM GRABBABLE FROM HERE
	return rNotifications
}

func MainSetup() {
	logging.Info("Setup Database")
	model.Setup(flags.ConfigFolder() + databaseFile)
	if _, err := os.Stat(flags.ConfigFolder() + tempSaveFile); !os.IsNotExist(err) {
		file, _ := os.Open(flags.ConfigFolder() + tempSaveFile)
		defer file.Close()
		json.NewDecoder(file).Decode(&server.Sessions)
		os.Remove(flags.ConfigFolder() + tempSaveFile)
	}
}

func MainShutdown() {
	if len(server.Sessions) == 0 {
		return
	}
	b, _ := json.Marshal(server.Sessions)
	reader := bytes.NewReader(b)
	file, _ := os.Create(flags.ConfigFolder() + tempSaveFile)
	defer file.Close()
	io.Copy(file, reader)
}

func GRPCSystems(grpcs *grpc.Server) {
	pbu.RegisterUserServer(grpcs, &server.UserServer{})
}
