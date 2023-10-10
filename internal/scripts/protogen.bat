cd C:\Users\user\go\src

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/common --go_opt=paths=source_relative common.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/internal --go_opt=paths=source_relative message.internal.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/protoc --go_opt=paths=source_relative core.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/protoc --go_opt=paths=source_relative core.response.proto

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/access --go_opt=paths=source_relative access/access.proto

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/game --go_opt=paths=source_relative game/data.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/game --go_opt=paths=source_relative game/game.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/game --go_opt=paths=source_relative game/game.response.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/game --go_opt=paths=source_relative game/game.notification.proto


protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/rank --go_opt=paths=source_relative rank/rank.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/rank --go_opt=paths=source_relative rank/rank.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/rank --go_opt=paths=source_relative rank/rank.response.proto


protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/social --go_opt=paths=source_relative social/social.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/social --go_opt=paths=source_relative social/social.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/social --go_opt=paths=source_relative social/social.response.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/social --go_opt=paths=source_relative social/social.notification.proto

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/challenge --go_opt=paths=source_relative challenge/challenge.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/challenge --go_opt=paths=source_relative challenge/challenge.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/challenge --go_opt=paths=source_relative challenge/challenge.response.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto/challenge --go_opt=paths=source_relative challenge/challenge.notification.proto

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto/services --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false service/access.service.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto/services --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false service/game.service.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto/services --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false service/social.service.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto/services --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false service/rank.service.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go-grpc_out=github.com/justjack1521/mevium/pkg/genproto/services --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false service/challenge.service.proto