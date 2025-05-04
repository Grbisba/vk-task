.PHONY: tag
tag:
	git tag protoc/proto/pubsub/v1.0.1 && ([ $$? -eq 0 ] && echo "success!") || echo "failure!"
	git tag subpub/v1.0.1 && ([ $$? -eq 0 ] && echo "success!") || echo "failure!"
	git tag vk-task/v1.0.0 && ([ $$? -eq 0 ] && echo "success!") || echo "failure!"

.PHONY: protogen
protogen:
	protoc ./proto/pubsub/subpub.proto --go_out=./protoc --go_opt=paths=source_relative --go-grpc_out=./protoc --go-grpc_opt=paths=source_relative
