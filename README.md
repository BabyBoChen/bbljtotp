# BBLJ TOTP
https://bbljtotp.onrender.com/

A TOTP Verifier App Utilizing github.com/pquerna/otp/totp Package

This TOTP generator is compatible with Google Authenticator (Google Play).

Scan the generated QR code with Google Authenticator app in order to save your secret key in your mobile device.

For server side verification, use this following URL with proper query parameters:

https://bbljtotp.onrender.com/verify?secret=${secret}&totp=${6-digits-passcode}
