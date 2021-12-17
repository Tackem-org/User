package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/registerService"
	pb "github.com/Tackem-org/Proto/pb/registration"
	"github.com/Tackem-org/User/flags"
	"github.com/Tackem-org/User/gprcServer"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/web"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	fmt.Println("Starting Tackem User System")

	if !helpers.InDockerCheck() {
		return
	}
	logging.Setup(*flags.LogFile, *flags.Verbose)
	logging.Info("Logger Started")

	registerService.Data = registerService.NewRegister()
	registerService.Data.Setup("user", "system", false, true, []*pb.NavItem{
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
	})

	logging.Info("Setup Registration Data")
	wg := &sync.WaitGroup{}

	logging.Info("Setup Database")
	model.Setup()

	logging.Info("Setup Web Service")
	web.Setup()

	logging.Info("Setup GPRC Service")
	gprcServer.Setup(wg)

	if !registerService.Data.Connect() {
		shutdown(wg, false)
	}
	logging.Info("Registration Done")
	captureInterupt(wg)
	wg.Wait()
	fmt.Println("Shutdown Complete Exiting Cleanly")
	os.Exit(0)
}

func shutdown(wg *sync.WaitGroup, registered bool) {

	if registered {
		registerService.Data.Disconnect()
		logging.Info("DeRegistration Done")
	}

	gprcServer.Shutdown(wg)
	logging.Info("Shutdown gRPC Server")

	web.Shutdown(wg)
	logging.Info("Shutdown Web Server")

	logging.Info("Closing Logger")
	logging.Shutdown()

}

func captureInterupt(wg *sync.WaitGroup) {
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	go func(wg *sync.WaitGroup) {
		<-termChan
		logging.Warning("SIGTERM received. Shutdown process initiated")
		shutdown(wg, true)
	}(wg)
}
