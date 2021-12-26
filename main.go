package main

import (
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

func main() {
	pflag.Parse()
	system.Run(system.SetupData{
		BaseData: system.BaseData{
			ServiceName: "user",
			ServiceType: "system",
			Multi:       false,
			WebAccess:   true,
			NavItems: []*pb.NavItem{
				{
					LinkType: pb.LinkType_User,
					Title:    "User",
					Icon:     "user",
					Path:     "user/",
				},
				{
					LinkType: pb.LinkType_Admin,
					Title:    "Users",
					Icon:     "users",
					Path:     "user/",
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
			system.WebAddPath("{{number:userid}}", web.UserIDPage)
			system.WebAddPath("{{string:username}}", web.UserNamePage)
		},
		MainSystem: program,
	})
}

func program() {
	logging.Info("Setup Database")
	model.Setup(*databaseFile)
}
