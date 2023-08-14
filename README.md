# bbljtotp
https://localhost:8080

A TOTP Verifier App Utilizing github.com/pquerna/otp/totp Package

This TOTP generator is compatible with Google Authenticator (Google Play).

Scan the generated QR code with Google Authenticator app in order to save your secret code in your mobile device.

For server side verification, use this following URL with proper query parameters:

https://localhost:8080/verify?secret=${secret}&totp=${6-digits-passcode})https://localhost:8080

A TOTP Verifier App Utilizing github.com/pquerna/otp/totp Package

This TOTP generator is compatible with Google Authenticator (Google Play).

Scan the generated QR code with Google Authenticator app in order to save your secret code in your mobile device.

For server side verification, use this following URL with proper query parameters:

https://localhost:8080/verify?secret=${secret}&totp=${6-digits-passcode}
