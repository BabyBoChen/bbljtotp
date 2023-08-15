package utils

import "github.com/pquerna/otp/totp"

func Verify(secret string, passcode string) bool {
	return totp.Validate(passcode, secret)
}
