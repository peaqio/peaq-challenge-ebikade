protoc -I=proto/rate -I=proto/third_party --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative  --go_out proto/rate --go-grpc_out proto/rate $(find proto/rate -iname "*.proto")
# *** Make sure "github.com/grpc-ecosystem/grpc-gateway/v2" is installed 
# and on go.mod file ***
protoc -I=proto/rate -I=proto/third_party --grpc-gateway_out=paths=source_relative,logtostderr=true:proto/rate $(find proto/rate -iname "*.proto")
# Swagger gen
protoc  --proto_path=proto/rate --proto_path=proto/third_party --openapiv2_opt logtostderr=true --openapiv2_out=allow_merge=true,merge_file_name=server.swagger:proto/swagger proto/rate/server.proto