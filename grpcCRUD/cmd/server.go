package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	v1 "grpcCRUD/api/service/v1"
	conf "grpcCRUD/conf"
	grpc "grpcCRUD/pkg/grpc"
	rest "grpcCRUD/pkg/rest"

	_ "github.com/go-sql-driver/mysql"
)

var cfg Config

type Config struct {
	GRPCPort            string
	HTTPPort            string
	DataStoreDBHost     string
	DataStoreDBUser     string
	DataStroeDBPassword string
	DataStoreDBSchema   string
}

//輸入參數 GPRC端口、DB地址、密碼、表
func init() {
	cfg = Config{}
	flag.StringVar(&cfg.GRPCPort, "grpc-port", conf.Port, "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", conf.HTTPPort, "Http port to bind")
	flag.StringVar(&cfg.DataStoreDBHost, "db-host", conf.DbHost, "db host")
	flag.StringVar(&cfg.DataStoreDBUser, "db-user", conf.DbUser, "db-user")
	flag.StringVar(&cfg.DataStroeDBPassword, "db-passward", conf.DbPassword, "db-password")
	flag.StringVar(&cfg.DataStoreDBSchema, "db-schema", conf.DbSchema, "db-schema")
	fmt.Println("init:" + cfg.GRPCPort)
	flag.Parse()
}

func RunServer() error {
	ctx := context.Background()
	// fmt.Println(len(cfg.GRPCPort))
	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server %s", cfg.GRPCPort)
	}
	fmt.Println(cfg.HTTPPort)
	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP server: %s", cfg.HTTPPort)
	}

	param := "parseTime=true"
	//連接資料庫字串
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DataStoreDBUser,
		cfg.DataStroeDBPassword,
		cfg.DataStoreDBHost,
		cfg.DataStoreDBSchema,
		param,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return fmt.Errorf("連接數據失敗: %v", err)
	}
	defer db.Close()

	v1API := v1.NewToDoServiceServer(db)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
