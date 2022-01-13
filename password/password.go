package password

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"os"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"

	"golang.org/x/crypto/pbkdf2"
)

var (
	salt []byte
)

const (
	saltFile   = "/config/salt.dat"
	saltLength = 16
	MaxLength  = 8
)

func SetupSalt() {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[password.SetupSalt()]")
	salt = make([]byte, saltLength)
	f, err := os.Open(saltFile)
	if err == nil {
		_, err = f.Read(salt)
		if err != nil {
			logging.Error("Error In Reading Salt File")
			createSaltFile()
		}
		f.Close()
	} else {
		createSaltFile()
	}
}

func createSaltFile() {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[password.createSaltFile()]")
	f, err := os.Create(saltFile)
	if err != nil {
		logging.Errorf("Error In Creating Salt File: %s", err.Error())
	}
	_, err = rand.Read(salt)
	if err != nil {
		logging.Errorf("Error In Generating Salt Bytes: %s", err.Error())
	}
	_, err = f.Write(salt)
	if err != nil {
		logging.Errorf("Error In Writing Salt File: %s", err.Error())
	}
	err = f.Close()
	if err != nil {
		logging.Errorf("Error In Closing Salt File: %s", err.Error())
	}
}

func Hash(password string) string {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[password.Hash(password string) string] {password=XXXXXXX}")
	return fmt.Sprintf("%x", pbkdf2.Key([]byte(password), []byte(salt), 4096, len(salt)*8, sha512.New))
}
