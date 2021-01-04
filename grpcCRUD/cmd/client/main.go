package main

import (
	"database/sql"
	"flag"
	"fmt"

	// v1 "grpcCRUD/api/v1"
	"log"
	"time"

	v1 "grpcCRUD/api/v1"
	conf "grpcCRUD/conf"

	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

func main() {
	// 之前剛看到這個錯誤指令的時候，一時間沒想到太多，┬但對照source code後發現在設定host的時候，沒設定好會出現：Error while dialing dial tcp: address 127.0.0.18999
	address := flag.String("server", "127.0.0.1:"+conf.Port, "gRPC server in format host:port")
	flag.Parse()
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client connect to  Server failed : %v", err)
	}
	defer conn.Close()
	c := v1.NewToDoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// CreateData(c, ctx)
	id, ctx := CreateData(c, ctx)
	ReadData(c, id, ctx)
	// ctx = ReadData(c, id, ctx)
	UpdateData(c, id, ctx)
	DeleteData(c, id, ctx)
	ReadAllData(c, ctx)
}

func CreateData(c v1.ToDoServiceClient, ctx context.Context) (int64, context.Context) {
	//定義上下文超時屬性
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	//轉換要插入的時間格式
	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	pfx := t.Format(time.RFC3339Nano)

	//調用遠端Server的Creat方法
	CreateReq := v1.CreateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Title:       "title(" + pfx + ")",
			Description: "description(" + pfx + ")",
			Reminder:    reminder,
		},
	}
	CreateRes, err := c.Create(ctx, &CreateReq)
	id := CreateRes.Id
	if err != nil {
		log.Fatalf("Connected to server, but failed to Create: %v", err)
	}
	fmt.Printf("Create success:%v", CreateRes)
	// id := CreateRes.Id
	// fmt.Sprintf("#{CreateRes}")
	return id, ctx
}

func ReadData(c v1.ToDoServiceClient, id int64, ctx context.Context) context.Context {
	//調用Server端的Read方法
	ReadReq := v1.ReadRequest{Api: apiVersion, Id: id}
	//定義上下文超時屬性
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	ReadRes, err := c.Read(ctx, &ReadReq)
	if err != nil {
		log.Fatalf("讀取格式失敗: %v", err)
	}
	log.Printf("Read success:", ReadRes)
	return ctx
}

func UpdateData(c v1.ToDoServiceClient, id int64, ctx context.Context) context.Context {
	//轉換要插入的時間格式
	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	//調用Server端的Update方法
	fmt.Println(id)
	UpdateReq := v1.UpdateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Id:          id,
			Title:       "Update req.ToDo.Title",
			Description: "Update req.ToDo.Description" + " updated",
			Reminder:    reminder,
		},
	}
	UpdateRes, err := c.Update(ctx, &UpdateReq)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update: %v", UpdateRes)
	return ctx
}

func DeleteData(c v1.ToDoServiceClient, id int64, ctx context.Context) context.Context {
	//調用Server端的Delete方法
	DeleteReq := v1.DeleteRequest{
		Api: apiVersion,
		Id:  id,
	}
	DeleteRes, err := c.Delete(ctx, &DeleteReq)

	if err != nil {
		log.Fatalf("Dalete failed: %v", err)
	}
	log.Printf("Deleted: %v", DeleteRes)
	return ctx
}

func ReadAllData(c v1.ToDoServiceClient, ctx context.Context) context.Context {
	//調用Server層的ReadAll方法
	ReadAllReq := v1.ReadAllRequest{
		Api: apiVersion,
	}
	// ReadAllRes, err := c.ReadAll(ctx, &ReadAllReq)
	ReadAllRes, err := c.ReadAll(ctx, &ReadAllReq)
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result %+v", ReadAllRes)
	return ctx
}

type ToDoServiceServer struct {
	db *sql.DB
}

func NewToDoServiceServer(db *sql.DB) *ToDoServiceServer {
	return &ToDoServiceServer{db: db}
}

//檢查api版本
func (s *ToDoServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			msg := "unsupported API version" + apiVersion
			return status.Error(codes.Unimplemented, msg)
		}
	}
	return nil
}
