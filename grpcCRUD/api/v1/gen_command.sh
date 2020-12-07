protoc todo-service.proto --go_out=plugins=grpc:.


protoc --proto_path=. --go_out=. *.proto


//protoc todo-service.proto --go_out=../../grpcCRUD/third_party/google/api/annotations.proto=../../grpcCRUD/third_party/////////////protoc-gen-swagger/options/annotations.proto,plugins=grpc:.