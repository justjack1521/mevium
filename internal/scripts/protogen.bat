cd C:\Users\user\go\src

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/common --go_opt=paths=source_relative common.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/internal --go_opt=paths=source_relative common.message.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/protoc --go_opt=paths=source_relative core.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/protoc --go_opt=paths=source_relative core.response.proto

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative access/access.proto

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protogame/data.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protogame/game.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protogame/game.response.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protogame/game.message.proto


protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protorank/rank.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protorank/rank.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protorank/rank.response.proto


protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protosocial/social.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protosocial/social.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protosocial/social.response.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protosocial/social.notification.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protosocial/social.message.proto

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protochallenge/challenge.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protochallenge/challenge.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protochallenge/challenge.response.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protochallenge/challenge.notification.proto

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false service/access.service.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false service/game.service.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false service/social.service.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false service/rank.service.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false service/challenge.service.proto