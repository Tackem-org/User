package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/structs"

	"github.com/Tackem-org/Global/system"
	pbc "github.com/Tackem-org/Proto/pb/config"
	pb "github.com/Tackem-org/Proto/pb/registration"
	pbuser "github.com/Tackem-org/Proto/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/server"
	"github.com/Tackem-org/User/static"
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
	system.Run(system.SetupData{
		BaseData: system.BaseData{
			ServiceName: "user",
			ServiceType: "system",
			Version: structs.Version{
				Major:  0,
				Minor:  0,
				Hotfix: 0,
			},
			Multi:     false,
			SingleRun: false,
			ConfigItems: []*pb.ConfigItem{
				{
					Key:          "user.password.minimum",
					DefaultValue: "8",
					Type:         pbc.ValueType_Uint,
					Label:        "Password Minimum Length",
					HelpText:     "what is the minimum password length",
					InputType:    pb.InputType_INumber,
					InputAttributes: &pb.InputAttributes{
						Other: []*pb.DictValue{
							{Name: "min", Value: "1"},
							{Name: "max", Value: "16"},
						},
					},
				},
			},
			NavItems: []*pb.NavItem{
				// {LinkType: pb.LinkType_User, Title: "User", Icon: "user", Path: "/"},
				{LinkType: pb.LinkType_User, Title: "Change Password", Icon: "user", Path: "/changepassword", Permission: "system_user_change_own_password"},
				{LinkType: pb.LinkType_User, Title: "Change Username", Icon: "user", Path: "/changeusername", Permission: "system_user_change_own_username"},
				{LinkType: pb.LinkType_User, Title: "Request New Username", Icon: "user", Path: "/requestusername", Permission: "system_user_request_change_of_username"},
				{LinkType: pb.LinkType_Admin, Title: "Users", Icon: "users", Path: "/", SubLinks: []*pb.NavItem{
					{LinkType: pb.LinkType_Admin, Title: "Groups", Icon: "user-shield", Path: "/groups"},
					{LinkType: pb.LinkType_Admin, Title: "Permissions", Icon: "key", Path: "/permissions"},
				},
				},
			},
		},
		LogFile:    *logFile,
		VerboseLog: *verbose,
		DebugLevel: debug.NONE,
		GPRCSystems: func(grpcs *grpc.Server) {
			pbuser.RegisterUserServer(grpcs, &server.UserServer{})
		},
		WebSystems: func() {
			system.WebSetup(&static.FS)
			system.WebAddAdminPath(&pb.AdminWebLinkItem{
				Path:        "/",
				PostAllowed: false,
				GetDisabled: false,
			}, admin.AdminRootPage)
			system.WebAddAdminPath(&pb.AdminWebLinkItem{
				Path:        "/edit/{{number:userid}}",
				PostAllowed: false,
				GetDisabled: false,
			}, admin.AdminUserIDPage)
			system.WebAddAdminPath(&pb.AdminWebLinkItem{
				Path:        "/groups",
				PostAllowed: false,
				GetDisabled: false,
			}, admin.AdminGroupsPage)
			// system.WebAddAdminWebSocket("/groups.ws", admin.AdminGroupsWebSocket)
			// system.WebAddAdminWebSocket("/edituser.ws", admin.AdminEditUserWebSocket)
			system.WebAddAdminPath(&pb.AdminWebLinkItem{
				Path:        "/permissions",
				PostAllowed: false,
				GetDisabled: false,
			}, admin.AdminPermissionsPage)
			system.WebAddPath(&pb.WebLinkItem{
				Path:        "/",
				Permission:  "",
				PostAllowed: false,
				GetDisabled: false,
			}, web.RootPage)
			system.WebAddPath(&pb.WebLinkItem{
				Path:        "/changepassword",
				Permission:  "system_user_change_own_password",
				PostAllowed: false,
				GetDisabled: false,
			}, web.ChangePasswordPage)
			system.WebAddPath(&pb.WebLinkItem{
				Path:        "/changeusername",
				Permission:  "system_user_change_own_username",
				PostAllowed: false,
				GetDisabled: false,
			}, web.ChangeUsernamePage)
			system.WebAddPath(&pb.WebLinkItem{
				Path:        "/requestusername",
				Permission:  "system_user_request_change_of_username",
				PostAllowed: false,
				GetDisabled: false,
			}, web.RequestUsernamePage)
			// system.WebAddPath("/edit", web.EditPage)
		},
		MainSetup: func() {
			logging.Info("Setup Database")
			model.Setup(*databaseFile)
			if _, err := os.Stat(tempSavePath); !os.IsNotExist(err) {
				loadData()
			}
		},
		Shutdown: func() {
			Save()
		},
	})
}

func Save() {

	if len(server.Sessions) == 0 {
		return
	}
	b, _ := json.Marshal(server.Sessions)
	reader := bytes.NewReader(b)
	file, _ := os.Create(tempSavePath)
	defer file.Close()
	io.Copy(file, reader)
}

func loadData() {
	file, _ := os.Open(tempSavePath)
	defer file.Close()
	json.NewDecoder(file).Decode(&server.Sessions)
	os.Remove(tempSavePath)
}
