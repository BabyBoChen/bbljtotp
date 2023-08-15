go env -w GOOS=js GOARCH=wasm
go build -o ../statics/totpWASM.wasm
