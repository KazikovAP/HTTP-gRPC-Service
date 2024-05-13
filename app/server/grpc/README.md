**Пример был взят с:** https://github.com/grpc/grpc-go/tree/master/examples/helloworld
**Quick start gRPC:** https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code
 
### Installing protoc on Windows
* Download `protoc-26.1-win64.zip` or `protoc-26.1-win32.zip` from https://github.com/protocolbuffers/protobuf/releases/
* Unzip and add location of the `protoc.exe` to your **PATH** environment variable
* Run `protoc --version` from command prompt to verify // *my version libprotoc 26.1*
* Verify the your `GOPATH` environment variable is set
* Run `go get -u github.com/golang/protobuf/protoc-gen-go` from command prompt. This should install the binary to `%GOPATH%/bin`
* Add `%GOPATH%/bin` to your **PATH** environment variable `export PATH="$PATH:$(go env GOPATH)/bin"`
* Open a new command prompt, navigate to your `.proto file`, run `protoc --go_out=. *.proto`

### Генерация кода:
* Командой `protoc --go_out=. *.proto` генерируется файл `grpc.pb.go`
* Командой `protoc --go_out=. --go-grpc_out=. grpc.proto` генерируется файл `grpc_grpc.pb.go`
* Подобной командой так же генерируется нужная обвязка для других поддерживаемых языков
* **go_out** означает что мы хотим сгенерировать код в этой папке для языка go
* **plugins=grpc** созначает что мы хотим использовать ещё плагин для генерации grpc-сервиса

### Запуск приложения
* Запуск сервера `go run server/server.go` из директории `HTTP-gRPC-Service/app/server/grpc`
* Запуск клиента `go run client/client.go` или с указанием имени `go run client/client.go --name=Aleksey` из директории `HTTP-gRPC-Service/app/server/grpc`
