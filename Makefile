.PHONY: tag
tag:
	git tag pubsub/v1.0.1 && ([ $$? -eq 0 ] && echo "success!") || echo "failure!"
	git tag subpub/v1.0.0 && ([ $$? -eq 0 ] && echo "success!") || echo "failure!"

.PHONY: protogen
protogen:
	protoc ./subpub.proto --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./protogen --go-grpc_opt=paths=source_relative
