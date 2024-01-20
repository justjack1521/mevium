cd %GOPATH%/src

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --csharp_out=github.com/justjack1521/mevium/internal/genproto/rank --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protorank/rank.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --csharp_out=github.com/justjack1521/mevium/internal/genproto/rank --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protorank/rank.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --csharp_out=github.com/justjack1521/mevium/internal/genproto/rank --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protorank/rank.response.proto