

cd pb

protoc -I . stream.proto --go_out=../proto --go_opt=paths=source_relative \
    --go-grpc_out=../proto --go-grpc_opt=paths=source_relative 