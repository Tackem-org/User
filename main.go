package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/setupData"
	"gorm.io/gorm/clause"

	"github.com/Tackem-org/Global/system"
	pbc "github.com/Tackem-org/Proto/pb/config"
	pb "github.com/Tackem-org/Proto/pb/registration"
	pbu "github.com/Tackem-org/Proto/pb/user"
	pbw "github.com/Tackem-org/Proto/pb/web"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/server"
	"github.com/Tackem-org/User/socket"
	"github.com/Tackem-org/User/socket/admin/editUser"
	"github.com/Tackem-org/User/socket/admin/group"
	"github.com/Tackem-org/User/static"
	"github.com/Tackem-org/User/tasks"
	"github.com/Tackem-org/User/web"
	"github.com/Tackem-org/User/web/admin"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

var (
	databaseFile = pflag.StringP("database", "d", "/config/User.db", "Database Location")
	logFile      = pflag.StringP("log", "l", "/logs/User.log", "Log Location")
	verbose      = pflag.BoolP("verbose", "v", false, "Outputs the log to the screen")
)

const (
	tempSavePath = "/config/tackemusersessionsdata.tmp"
)

func main() {
	pflag.Parse()
	system.Run(&setupData.SetupData{

		ServiceName: "user",
		ServiceType: "system",
		Version: structs.Version{
			Major:  0,
			Minor:  0,
			Hotfix: 0,
		},
		Multi:       false,
		SingleRun:   false,
		StartActive: true,
		ConfigItems: []*pb.ConfigItem{
			{
				Key:          "user.password.minimum",
				DefaultValue: "8",
				Type:         pbc.ValueType_Uint,
				Label:        "Password Minimum Length",
				HelpText:     "what is the minimum password length",
				InputType:    pb.InputType_INumber,
				InputAttributes: &pb.InputAttributes{
					Other: map[string]string{
						"min": "1",
						"max": "16",
					},
				},
			},
		},
		NavItems: []*pb.NavItem{
			{LinkType: pb.LinkType_User, Title: "Change Password", Icon: "user", Path: "/changepassword", Permission: "system_user_change_own_password"},
			{LinkType: pb.LinkType_User, Title: "Change Username", Icon: "user", Path: "/changeusername", Permission: "system_user_change_own_username"},
			{LinkType: pb.LinkType_User, Title: "Request New Username", Icon: "user", Path: "/requestusername", Permission: "system_user_request_change_of_username"},
			{LinkType: pb.LinkType_Admin, Title: "Users", Icon: "users", Path: "/", SubLinks: []*pb.NavItem{
				{LinkType: pb.LinkType_Admin, Title: "Groups", Icon: "user-shield", Path: "/groups"},
				{LinkType: pb.LinkType_Admin, Title: "Permissions", Icon: "key", Path: "/permissions"},
			},
			},
		},
		MasterConf: "/config/user.json",
		LogFile:    *logFile,
		VerboseLog: *verbose,
		GRPCSystems: func(grpcs *grpc.Server) {
			pbu.RegisterUserServer(grpcs, &server.UserServer{})
		},

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
		TaskGrabber: func() []*pbw.TaskMessage {
			var rTasks []*pbw.TaskMessage
			var uChanges []model.UsernameRequest
			model.DB.Preload(clause.Associations).Find(&uChanges)
			for _, u := range uChanges {
				rTasks = append(rTasks, tasks.UserNameChangeRequest(&u))
			}
			return rTasks
		},
		MainSetup: func() {
			logging.Info("Setup Database")
			model.Setup(*databaseFile)
			if _, err := os.Stat(tempSavePath); !os.IsNotExist(err) {
				file, _ := os.Open(tempSavePath)
				defer file.Close()
				json.NewDecoder(file).Decode(&server.Sessions)
				os.Remove(tempSavePath)
			}
		},
		MainShutdown: func() {
			if len(server.Sessions) == 0 {
				return
			}
			b, _ := json.Marshal(server.Sessions)
			reader := bytes.NewReader(b)
			file, _ := os.Create(tempSavePath)
			defer file.Close()
			io.Copy(file, reader)
		},
	})
}
