## build grpc

<!-- 
protoc -I$@/${PROTO_DIR} \
--go_opt=module=${PACKAGE} \
--go_out=. \
--go-grpc_opt=module=${PACKAGE} \
--go-grpc_out=. \
$@/${PROTO_DIR}/*.proto


protoc -Igreet/proto \
--go_opt=module=github.com/achmadAlli/go-grpc \
--go_out=. \
--go-grpc_opt=module=github.com/achmadAlli/go-grpc \
--go-grpc_out=. \
greet/proto/*.proto

protoc -Igreet/proto \
--go_out=. \
--go_opt=module=github.com/achmadAlli/go-grpc \
--go-grpc_out=proto/build \
--go-grpc_opt=module=github.com/achmadAlli/go-grpc/proto/build \
proto/greet.proto