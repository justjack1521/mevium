cd C:\Users\Henry\go\src

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/common --go_opt=paths=source_relative common.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/services --go_opt=paths=source_relative access.proto

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/protoc --go_opt=paths=source_relative web.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/protoc --go_opt=paths=source_relative web.response.proto

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/game --go_opt=paths=source_relative game.proto

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto/services --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false access.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto/services --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false web.service.proto