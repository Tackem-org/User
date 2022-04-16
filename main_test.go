package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Tackem-org/Global/logging"
	pbw "github.com/Tackem-org/Global/pb/web"
	"github.com/Tackem-org/Global/system/setupData"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/server"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type MockLogging struct{}

func (l *MockLogging) Setup(logFile string, verbose bool)                          {}
func (l *MockLogging) Shutdown()                                                   {}
func (l *MockLogging) CustomLogger(prefix string) *log.Logger                      { return log.New(nil, prefix+": ", 0) }
func (l *MockLogging) Custom(prefix string, message string, values ...interface{}) {}
func (l *MockLogging) Info(message string, values ...interface{})                  {}
func (l *MockLogging) Warning(message string, values ...interface{})               {}
func (l *MockLogging) Error(message string, values ...interface{})                 {}
func (l *MockLogging) Todo(message string, values ...interface{})                  {}
func (l *MockLogging) Fatal(message string, values ...interface{}) error {
	return fmt.Errorf(message, values...)
}

//TODO FINISH THIS FUNC AS FAILS
func Test_main(t *testing.T) {
	sd = &setupData.SetupData{}
	assert.Panics(t, func() {
		main()
	})
}

func TestTaskGrabber(t *testing.T) {
	assert.NotPanics(t, func() { model.Setup("testTaskGrabber.db") })
	defer os.Remove("testTaskGrabber.db")

	ur1 := &model.UsernameRequest{
		RequestUserID: 2,
		Name:          "new",
	}
	model.DB.Create(ur1)
	r := TaskGrabber()
	assert.IsType(t, []*pbw.TaskMessage{}, r)
	assert.Len(t, r, 1)
}

func TestMainSetupAndShutdown(t *testing.T) {
	logging.I = &MockLogging{}
	pflag.Set("config", "")
	assert.NotPanics(t, func() {
		MainShutdown()
	})
	server.Sessions = []server.Session{
		{
			UserID:       1,
			SessionToken: "test",
			IPAddress:    "127.0.0.1",
			ExpireTime:   time.Now().Add(time.Second),
		},
	}

	assert.NotPanics(t, func() {
		MainShutdown()
		MainSetup()
	})
	os.Remove("Salt.dat")
	os.Remove("User.db")

}

func TestGRPCSystems(t *testing.T) {
	assert.NotPanics(t, func() {
		GRPCSystems(grpc.NewServer())
	})
}
