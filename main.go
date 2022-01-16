package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"

	"github.com/Tackem-org/Global/system"
	pb "github.com/Tackem-org/Proto/pb/registration"
	pbuser "github.com/Tackem-org/Proto/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/server"
	"github.com/Tackem-org/User/static"
	"github.com/Tackem-org/User/web"
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
			WebAccess: true,
			NavItems: []*pb.NavItem{
				{LinkType: pb.LinkType_User, Title: "User", Icon: "user", Path: "/"},
				{LinkType: pb.LinkType_Admin, Title: "Users", Icon: "users", Path: "/"},
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
			system.WebAddAdminPath("/", web.AdminRootPage)
			system.WebAddAdminPath("/{{number:userid}}", web.AdminUserIDPage)
			system.WebAddPath("/", web.RootPage)
			system.WebAddPath("/edit", web.EditPage)
			system.WebAddPath("/view/{{number:userid}}", web.UserIDPage)
			system.WebAddPath("/view/{{string:username}}", web.UserNamePage)
		},
		MainSetup: func() {
			logging.Info("Setup Database")
			model.Setup(*databaseFile)
			if _, err := os.Stat(tempSavePath); !os.IsNotExist(err) {
				loadData()
			}
		},
		Shutdown: func() {
			saveData()
		},
	})
}

func saveData() {

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
