cd %GOPATH%/src
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protoaccess/access.proto