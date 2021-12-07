package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Tackem-org/Global/indocker"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/User/flags"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	fmt.Println("Starting Tackem User System")

	if !indocker.Check() {
		return
	}
	logging.Setup(*flags.LogFile, *flags.Verbose)
	logging.Info("Logger Started")

	// wg := &sync.WaitGroup{}
	// wg.Add(2)
	// gprcServer.Run(wg)
	// //ADD IN OTHER SUB SYSTEMS HERE WHEN CODED
	// web.Run(wg)
	// captureInterupt(wg)

	// wg.Wait()
}

func shutdown(wg *sync.WaitGroup) {
	// gprcServer.Shutdown(wg)
	// logging.Info("Shutdown gRPC Server")

	// logging.Info("Shutdown Web Server")
	// config.Shutdown()
	// logging.Info("Closed Config")
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
