package password

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"os"

	"github.com/Tackem-org/Global/logging"

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
		f, err := os.Create(saltFile)
		if err != nil {
			logging.Error("Error In Creating Salt File:" + err.Error())
		}
		_, err = rand.Read(salt)
		if err != nil {
			logging.Error("Error In Generating Salt Bytes:" + err.Error())
		}
		_, err = f.Write(salt)
		if err != nil {
			logging.Error("Error In Writing Salt File:" + err.Error())
		}
		err = f.Close()
		if err != nil {
			logging.Error("Error In Closing Salt File:" + err.Error())
		}
	}
}

func Hash(password string) string {
	return fmt.Sprintf("%x", pbkdf2.Key([]byte(password), []byte(salt), 4096, len(salt)*8, sha512.New))
}
