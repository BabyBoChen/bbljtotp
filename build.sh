cd wasm
go env -w GOOS=js GOARCH=wasm
go build -o ../statics/totpWASM.wasm
cd ..
go env -w GOOS=linux GOARCH=amd64
go build -o bbljtotp.exe .