package password

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"os"

	"github.com/Tackem-org/Global/flags"
	"github.com/Tackem-org/Global/logging"

	"golang.org/x/crypto/pbkdf2"
)

var (
	Salt           []byte
	SaltFile       = "Salt.dat"
	CreateSaltFile = createSaltFile
)

const (
	saltLength = 16
	MaxLength  = 8
)

func SetupSalt() error {
	Salt = make([]byte, saltLength)
	f, err := os.Open(flags.ConfigFolder() + SaltFile)
	if err != nil {
		return CreateSaltFile()
	}

	f.Read(Salt)
	return f.Close()
}

func createSaltFile() error {
	f, err := os.Create(flags.ConfigFolder() + SaltFile)
	if err != nil {
		logging.Error("Error In Creating Salt File: %s", err.Error())
		return err
	}
	rand.Read(Salt)
	f.Write(Salt)
	f.Close()
	return nil
}

func Hash(password string) string {
	return fmt.Sprintf("%x", pbkdf2.Key([]byte(password), []byte(Salt), 4096, len(Salt)*8, sha512.New))
}
