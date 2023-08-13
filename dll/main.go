package main

import "C"
import (
	"github.com/pquerna/otp/totp"
)

//export Verify
func Verify(secret *C.char, passcode *C.char) C.int {
	totpSecret := C.GoString(secret)
	totpPasscode := C.GoString(passcode)
	isValid := totp.Validate(totpPasscode, totpSecret)
	if isValid {
		return C.int(1)
	} else {
		return C.int(0)
	}
}

func main() {

}
