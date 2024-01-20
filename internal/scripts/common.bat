cd %GOPATH%/src

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --csharp_out=github.com/justjack1521/mevium/internal/genproto --go_opt=paths=source_relative protocommon/common.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protocommon/common.message.proto