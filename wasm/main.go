package main

import (
	"fmt"
	"syscall/js"
	"time"

	"github.com/pquerna/otp/totp"
)

var generateTOTP js.Func = js.FuncOf(_generateTOTP)

func _generateTOTP(this js.Value, args []js.Value) interface{} {

	err := "unknown"
	hasError := false
	var secret js.Value = js.ValueOf("")
	var passcode_js js.Value

	if len(args) == 1 {
		secret = args[0]
	} else {
		err = "this method takes one arguments (string: secret)"
		hasError = true
	}

	if !hasError {
		if secret.Type() != js.TypeString {
			err = "the first parameter should be a string"
			hasError = true
		}
	}

	if !hasError {
		passcode, err2 := totp.GenerateCode(secret.String(), time.Now().UTC())
		if err2 == nil {
			passcode_js = js.ValueOf(passcode)
		} else {
			err = err2.Error()
			hasError = true
		}
	}

	if !hasError {
		return passcode_js
	} else {
		return js.Error{
			Value: js.ValueOf(err),
		}
	}
}

func main() {
	js.Global().Set("generateTOTP", generateTOTP)
	fmt.Println("totpWASM loaded")
	<-make(chan bool)
}
