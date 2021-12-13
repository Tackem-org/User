package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Tackem-org/Global/global"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/registerService"
	pb "github.com/Tackem-org/Proto/pb/registration"
	"github.com/Tackem-org/User/flags"
	"github.com/Tackem-org/User/gprcServer"
	"github.com/Tackem-org/User/web"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	fmt.Println("Starting Tackem User System")

	if !global.InDockerCheck() {
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

	logging.Info("Setup Web System Data")
	web.Setup()
	gprcServer.Run(wg)

	if !registerService.Data.Connect() {
		shutdown(wg)
	}
	logging.Info("Registration Done")
	captureInterupt(wg)
	wg.Wait()
}

func shutdown(wg *sync.WaitGroup) {
	registerService.Data.Disconnect()
	logging.Info("DeRegistration Done")
	gprcServer.Shutdown(wg)
	logging.Info("Shutdown gRPC Server")

	logging.Info("Closing Logger")
	logging.Shutdown()
	fmt.Println("Shutdown Complete Exiting Cleanly")
	os.Exit(0)
}

func captureInterupt(wg *sync.WaitGroup) {
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	go func(wg *sync.WaitGroup) {
		<-termChan
		logging.Warning("SIGTERM received. Shutdown process initiated\n")
		shutdown(wg)
	}(wg)
}
