## build grpc
protoc \
--go_out=. \
--go_opt=module=github.com/achmadAlli/go-grpc \
--go-grpc_out=proto/build/ \
--go-grpc_opt=module=github.com/achmadAlli/go-grpc/proto/build \
proto/greet.proto