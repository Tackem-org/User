package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/Tackem-org/Global/logging"

	"github.com/Tackem-org/Global/system"
	pb "github.com/Tackem-org/Proto/pb/registration"
	pbuser "github.com/Tackem-org/Proto/pb/user"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/static"
	"github.com/Tackem-org/User/userServer"
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
			Multi:       false,
			SingleRun:   false,
			WebAccess:   true,
			NavItems: []*pb.NavItem{
				{
					LinkType: pb.LinkType_User,
					Title:    "User",
					Icon:     "user",
					Path:     "/",
				},
				{
					LinkType: pb.LinkType_Admin,
					Title:    "Users",
					Icon:     "users",
					Path:     "/",
				},
			},
		},
		LogFile:    *logFile,
		VerboseLog: *verbose,
		GPRCSystems: func(server *grpc.Server) {
			pbuser.RegisterUserServer(server, userServer.NewUserServer())
		},
		WebSystems: func() {
			system.WebSetup(&static.FS)
			system.WebAddPath("/", web.RootPage)
			system.WebAddAdminPath("/", web.AdminRootPage)
			system.WebAddAdminPath("/{{number:userid}}", web.AdminUserIDPage)
			system.WebAddPath("/{{number:userid}}", web.UserIDPage)
			system.WebAddPath("/{{string:username}}", web.UserNamePage)
		},
		MainSystem: func() {
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

	if len(userServer.Sessions) == 0 {
		return
	}
	b, _ := json.Marshal(userServer.Sessions)
	reader := bytes.NewReader(b)
	file, _ := os.Create(tempSavePath)
	defer file.Close()
	io.Copy(file, reader)
}

func loadData() {
	file, _ := os.Open(tempSavePath)
	defer file.Close()
	json.NewDecoder(file).Decode(&userServer.Sessions)
	os.Remove(tempSavePath)
}
