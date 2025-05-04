module github.com/Grbisba/vk-task/client

go 1.24.2

replace github.com/Grbisba/vk-task/protoc/proto/pubsub => ../protoc/proto/pubsub

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/heetch/confita v0.10.0
	github.com/pkg/errors v0.9.1
	go.uber.org/fx v1.23.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.72.0
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/Grbisba/vk-task/subpub v1.0.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	go.uber.org/dig v1.18.1 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250428153025-10db94c68c34 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
