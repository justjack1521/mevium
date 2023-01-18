cd C:\Users\Henry\go\src

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/common --go_opt=paths=source_relative common.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/services --go_opt=paths=source_relative access.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto/services --go-grpc_opt=paths=source_relative access.proto