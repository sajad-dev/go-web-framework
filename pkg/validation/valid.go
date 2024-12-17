package validation

import (
	"fmt"
	"net"
	"path/filepath"
	"regexp"
	"strconv"
	"unicode"
)



type Blob struct {
	Data     []byte 
	MIMEType string 
	FileName string 
}

func Required(str any) string {
	if str == nil || str == "" {
		return Message["required"]
	}
	return ""
}

func Accepted(value bool) string {
	if !value {
		return Message["accepted"]
	}
	return ""
}

func ActiveURL(url string) string {
	if !isValidURL(url) {
		return Message["active_url"]
	}
	return ""
}

func Alpha(str string) string {
	for _, r := range str {
		if !unicode.IsLetter(r) {
			return Message["alpha"]
		}
	}
	return ""
}

func AlphaNumeric(str string) string {
	for _, r := range str {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return Message["alpha_numeric"]
		}
	}
	return ""
}

func Email(email string) string {
	if !isValidEmail(email) {
		return Message["email"]
	}
	return ""
}

func Numeric(value string) string {
	if _, err := strconv.ParseFloat(value, 64); err != nil {
		return Message["numeric"]
	}
	return ""
}

func IP(ip string) string {
	if net.ParseIP(ip) == nil {
		return Message["ip"]
	}
	return ""
}

func Boolean(value any) string {
	switch v := value.(type) {
	case bool:
		return ""
	case string:
		if v == "true" || v == "false" {
			return ""
		}
	}
	return Message["boolean"]
}

func URL(url string) string {
	if !isValidURL(url) {
		return Message["url"]
	}
	return ""
}

// Min و Max برای طول رشته‌ها
func Min(str string, minLength int) string {
	if len(str) < minLength {
		return fmt.Sprintf(Message["min"], minLength)
	}
	return ""
}

func Max(str string, maxLength int) string {
	if len(str) > maxLength {
		return fmt.Sprintf(Message["max"], maxLength)
	}
	return ""
}

// چک کردن شماره تلفن
func Phone(phone string) string {
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	if !re.MatchString(phone) {
		return Message["phone"]
	}
	return ""
}

func Confirmed(password, confirmedPassword string) string {
	if password != confirmedPassword {
		return Message["confirmed"]
	}
	return ""
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func isValidURL(url string) bool {
	re := regexp.MustCompile(`^(http|https):\/\/[a-z0-9.-]+$`)
	return re.MatchString(url)
}

func File(blob *Blob) string {
	if blob == nil || len(blob.Data) == 0 {
		return Message["file_required"]
	}
	return ""
}

func MIMEType(blob *Blob, allowedTypes []string) string {
	for _, mimeType := range allowedTypes {
		if blob.MIMEType == mimeType {
			return ""
		}
	}
	return fmt.Sprintf(Message["mimes"], allowedTypes)
}

func MaxSize(blob *Blob, maxSize int) string {
	if len(blob.Data) > maxSize {
		return fmt.Sprintf(Message["max_size"], maxSize)
	}
	return ""
}

func Extension(blob *Blob, allowedExtensions []string) string {
	ext := filepath.Ext(blob.FileName)
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return ""
		}
	}
	return fmt.Sprintf(Message["extension"], allowedExtensions)
}
