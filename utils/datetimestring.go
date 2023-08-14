package utils

import (
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
)

func DateTimeString() string {
	return strings.Replace(time.Now().UTC().Format("20060102150405.000"), ".", "", 1)
}

func Verify(secret string, passcode string) bool {
	return totp.Validate(passcode, secret)
}
