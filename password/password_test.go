package password_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/User/password"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

type MockLogging struct {
	ErrorCount int
	ErrorMsg   string
}

func (l *MockLogging) Setup(logFile string, verbose bool)                          {}
func (l *MockLogging) Shutdown()                                                   {}
func (l *MockLogging) CustomLogger(prefix string) *log.Logger                      { return log.New(nil, prefix+": ", 0) }
func (l *MockLogging) Custom(prefix string, message string, values ...interface{}) {}
func (l *MockLogging) Info(message string, values ...interface{})                  {}
func (l *MockLogging) Warning(message string, values ...interface{})               {}
func (l *MockLogging) Error(message string, values ...interface{}) {
	l.ErrorCount++
	l.ErrorMsg = fmt.Sprintf(message, values...)
}
func (l *MockLogging) Todo(message string, values ...interface{}) {}
func (l *MockLogging) Fatal(message string, values ...interface{}) error {
	return fmt.Errorf(message, values...)
}

func TestSetupSalt(t *testing.T) {
	pflag.Set("config", "")
	assert.Nil(t, password.SetupSalt())
	assert.Nil(t, password.SetupSalt())
	os.Remove(password.SaltFile)
}

func TestCreateSaltFile(t *testing.T) {
	l := &MockLogging{}
	logging.I = l
	pflag.Set("config", "/")
	assert.Error(t, password.CreateSaltFile())
	assert.Equal(t, 1, l.ErrorCount)
	pflag.Set("config", "")
	assert.Nil(t, password.CreateSaltFile())
	os.Remove(password.SaltFile)
}

func TestHash(t *testing.T) {
	password.Salt = []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
	hash := "1d181c83f2c53e8003adf3dfba066845c247f90e87bccdab1e61c92f60925889137b7d0c5e4dabb25a7705ec8513f6835144fb9a6a734b321e1d9be1f3bc792a37be6bbf6d5d2158ad584c18128aeeb03b3a4bdd5b09373ed4dc5b70286409ce57a1e2264e32e86607637beb36cfeb95ee697c95c0b28204a593769e067305d3"
	assert.Equal(t, hash, password.Hash("test"))
}
