package main

import (
	"flag"
	v1 "grpcCRUD/api/v1"
	"log"
	"time"

	conf "grpcCRUD/conf"

	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	address := flag.String("server", "127.0.0.1"+conf.Port, "gRPC server in format host:port")
	flag.Parse()
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Client 連接 Server 失敗: %v", err)
	}
	defer conn.Close()
	c := v1.NewToDoServiceClient(conn)
	id := CreateData(c).GetId()
	ReadData(c, id)
	UpdateData(c, id)
	DeleteData(c, id)
	ReadAllData(c)
}

func CreateData(c v1.ToDoServiceClient) *v1.CreateResponse {
	apiVersion := "v1"
	//定義上下文超時屬性
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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
	if err != nil {
		log.Fatalf("Connected to server, but failed to Create: %v", err)
	}
	log.Printf("Create success:%v", CreateRes)
	// id := CreateRes.Id
	// fmt.Sprintf("#{CreateRes}")
	return CreateRes
}

func ReadData(c v1.ToDoServiceClient, id int64) {
	apiVersion := "v1"
	//調用Server端的Read方法
	ReadReq := v1.ReadRequest{Api: apiVersion, Id: id}
	//定義上下文超時屬性
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ReadRes, err := c.Read(ctx, &ReadReq)
	if err != nil {
		log.Fatalf("讀取格式失敗: %v", err)
	}
	log.Printf("Read success:", ReadRes)
}

func UpdateData(c v1.ToDoServiceClient, id int64) {
	apiVersion := "v1"
	//定義上下文超時屬性
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//轉換要插入的時間格式
	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	//調用Server端的Update方法
	UpdateReq := v1.UpdateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Id:          id,
			Title:       "Update Title",
			Description: "update description",
			Reminder:    reminder,
		},
	}
	UpdateRes, err := c.Update(ctx, &UpdateReq)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update: %v", UpdateRes)
}

func DeleteData(c v1.ToDoServiceClient, id int64) {
	apiVersion := "v1"
	//超時標準
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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
}

func ReadAllData(c v1.ToDoServiceClient) {
	apiVersion := "v1"
	//超時標準
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//調用Server端的ReadAll方法
	ReadAllReq := v1.ReadAllRequest{
		Api: apiVersion,
	}
	ReadAllRes, err := c.ReadAll(ctx, &ReadAllReq)
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result %v", ReadAllRes)
}
