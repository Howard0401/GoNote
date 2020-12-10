#一般匯出grpc
protoc --proto_path=. --go_out=. --go-grpc_out=require_unimplemented_servers=false:. *.proto
#匯出gateway
protoc --proto_path=. --grpc-gateway_out=logtostderr=true:. todo-service.proto


protoc --proto_path=. --grpc-gateway_out=. --grpc-gateway_opt=logtostderr=true  --grpc-gateway_opt generate_unbound_methods=true todo-service.proto
--grpc-gateway_opt
#匯出Swagger(JSON)
protoc --proto_path=. --swagger_out=logtostderr=true:. todo-service.proto


#好像不能用了
protoc todo-service.proto --go_out=plugins=grpc:.
#會有一個unimplemented，目前還沒探討怎麼實作，所以先設置為false
protoc --proto_path=. --go_out=. --go-grpc_out=. *.proto


protoc --proto_path=. --go_out=. --go-grpc_out=require_unimplemented_servers=false:. todo-service.proto


protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 todo-service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1 todo-service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 todo-service.proto