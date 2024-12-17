package testutils

import (
	"fmt"
	"math/rand"
	"time"
)

type TestFunc func() (string, string)
type TestFuncWithArgs func(int) (string, string)

var TestMap = map[string]TestFunc{
	"required":     TestRequired,
	"accepted":     TestAccepted,
	"activeURL":    TestActiveURL,
	"alpha":        TestAlpha,
	"alphaNumeric": TestAlphaNumeric,
	"email":        TestEmail,
	"numeric":      TestNumeric,
	"ip":           TestIP,
	"boolean":      TestBoolean,
	"url":          TestURL,
	"phone":        TestPhone,
	"confirmed":    TestConfirmed,
}

var TestMapWithArgs = map[string]TestFuncWithArgs{
	"min": TestMin,
	"max": TestMax,
}

type Blob struct {
	Data     []byte
	MIMEType string
	FileName string
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("0123456789")
var specialChars = []rune("!@#$%^&*()")

// Generate a random string of a specified length
func randomString(length int) string {
	b := make([]rune, length)
	rand.Seed(time.Now().UnixNano()) 
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Generate a random alphanumeric string of a specified length
func randomAlphaNumeric(length int) string {
	b := make([]rune, length)
	rand.Seed(time.Now().UnixNano()) 

	for i := range b {
		if rand.Intn(2) == 0 {
			b[i] = letters[rand.Intn(len(letters))]
		} else {
			b[i] = numbers[rand.Intn(len(numbers))]
		}
	}
	return string(b)
}

// Generate a random numeric string of a specified length
func randomNumeric(length int) string {
	b := make([]rune, length)
	rand.Seed(time.Now().UnixNano()) 

	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(b)
}

// Generate a random email
func randomEmail() string {

	return randomString(5) + "@example.com"
}

// Generate a random IP address
func randomIP() string {

	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// Generate a random phone number
func randomPhone() string {

	return fmt.Sprintf("+%s", randomNumeric(10))
}

// Generate a random URL
func randomURL() string {

	return fmt.Sprintf("http://%s.com", randomString(8))
}

// Generate random test data
func TestRequired() (string, string) {
	return "", "  " // valid: empty string; invalid: whitespace
}

func TestAccepted() (string, string) {
	return "true", "false" // valid: accepted; invalid: not accepted
}

func TestActiveURL() (string, string) {
	return randomURL(), "invalid_url" // valid: random URL; invalid: not a URL
}

func TestAlpha() (string, string) {
	return randomString(10), randomAlphaNumeric(10) // valid: only letters; invalid: alphanumeric
}

func TestAlphaNumeric() (string, string) {
	return randomAlphaNumeric(10), randomString(10) // valid: alphanumeric; invalid: only letters
}

func TestEmail() (string, string) {
	return randomEmail(), "invalid_email" // valid: random email; invalid: incorrect format
}

func TestNumeric() (string, string) {
	return randomNumeric(5), randomString(5) // valid: numeric string; invalid: non-numeric
}

func TestIP() (string, string) {
	return randomIP(), "invalid_ip" // valid: random IP; invalid: not an IP
}

func TestBoolean() (string, string) {
	return "true", "not_a_boolean" // valid: boolean; invalid: non-boolean string
}

func TestURL() (string, string) {
	return randomURL(), "not_a_url" // valid: random URL; invalid: not a URL
}

func TestMin(num int) (string, string) {
	return randomString(num+1), randomString(num-1) // valid: meets min length; invalid: below min length
}

func TestMax(num int) (string, string) {
	return randomString(2000 % num), randomString(num * 2) // valid: within max length; invalid: exceeds max length
}

func TestPhone() (string, string) {
	return randomPhone(), "12345" // valid: proper phone; invalid: too short
}

func TestConfirmed() (string, string) {
	return randomString(8), randomString(8) // valid: passwords match; invalid: passwords don't match
}
