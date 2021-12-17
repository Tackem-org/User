package gprcServer

import (
	"fmt"
	"net"
	"sync"

	"github.com/Tackem-org/Global/checkerServer"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/registerService"
	"github.com/Tackem-org/Global/remoteWebSystem"
	pbchecker "github.com/Tackem-org/Proto/pb/checker"
	pbremoteweb "github.com/Tackem-org/Proto/pb/remoteweb"
	pbuser "github.com/Tackem-org/Proto/pb/user"
	"github.com/Tackem-org/User/userServer"
	"google.golang.org/grpc"
)

var (
	server *grpc.Server
)

func Setup(wg *sync.WaitGroup) {
	server = grpc.NewServer()
	registerSystems()

	wg.Add(1)
	go func() {
		port := fmt.Sprint(registerService.Data.GetPort())
		listen, err := net.Listen("tcp", ":"+port)
		if err != nil {
			logging.Error("gPRC could not listen on port " + port)
			logging.Fatal(err)
		}
		if err := server.Serve(listen); err != nil {
			logging.Fatal(err)
		}
	}()
	logging.Info("Starting gRPC server")

}

func Shutdown(wg *sync.WaitGroup) {
	server.Stop()
	wg.Done()
}

func registerSystems() {
	pbremoteweb.RegisterRemoteWebServer(server, remoteWebSystem.NewServer())
	pbchecker.RegisterCheckerServer(server, checkerServer.NewServer())

	//add services here
	pbuser.RegisterUserServer(server, userServer.NewUserServer())
}
