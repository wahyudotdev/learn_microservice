echo "Generate go file from proto"
protoc --go_out=. ../proto/product/*.proto --proto_path=../proto
protoc --go-grpc_out=. ../proto/product/*.proto --proto_path=../proto