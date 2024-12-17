package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"os"
	"regexp"

	"github.com/sajad-dev/go-web-framework/pkg/exception"
	"github.com/sajad-dev/go-web-framework/pkg/helpers"
)

func IsValid(str string) bool {
	validate := regexp.MustCompile("^[a-zA-Z0-9_]+$")
	if !validate.MatchString(str) {
		if !helpers.IfThenElse(os.Getenv("DEBUG") == "true", true, false).(bool) {
			log.Panicln("Not a valid string")
		}
		return false
	}
	return true
}

func GenerateToken() string {
	const LENGTH int = 30

	randomBytes := make([]byte, 30)
	_, err := rand.Read(randomBytes)
	exception.Log(err)
	output := base64.StdEncoding.EncodeToString(randomBytes)
	return output
}

func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	pass := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(pass)
}

