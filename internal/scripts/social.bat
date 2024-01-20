cd %GOPATH%/src

protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --csharp_out=github.com/justjack1521/mevium/internal/genproto/social --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protosocial/social.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --csharp_out=github.com/justjack1521/mevium/internal/genproto/social --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protosocial/social.request.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --csharp_out=github.com/justjack1521/mevium/internal/genproto/social --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protosocial/social.response.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --csharp_out=github.com/justjack1521/mevium/internal/genproto/social --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protosocial/social.notification.proto
protoc --proto_path=github.com/justjack1521/mevium/internal/api/protobuf --go_out=github.com/justjack1521/mevium/pkg/genproto --go_opt=paths=source_relative protosocial/social.message.proto

