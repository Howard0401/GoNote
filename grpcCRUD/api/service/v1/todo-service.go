package v1

import (
	"context"
	"database/sql"
	"fmt"
	v1 "grpcCRUD/api/v1"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

//初始化db

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

func (s *ToDoServiceServer) mustEmbedUnimplementedToDoServiceServer() {
}

//連接資料庫
func (s *ToDoServiceServer) Connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "連接資料庫失敗"+err.Error())
	}
	return c, nil
}

func (s *ToDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	//檢查是否API符合規則
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := s.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close() //執行完Create後把Conn關閉，避免耗費資源
	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "參數錯誤："+err.Error())
	}
	//插入數據 (?,?,?) 避免sql injection
	res, err := c.ExecContext(ctx, "INSERT INTO ToDo(`Title`,`Description`,`Reminder`) VALUES(?,?,?)", req.ToDo.Title, req.ToDo.Description, reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "插入數據失敗:"+err.Error())
	}
	id, err := res.LastInsertId() //返回最新的Id
	if err != nil {
		return nil, status.Error(codes.Unknown, "建立id失敗:"+err.Error())
	}
	return &v1.CreateResponse{Api: apiVersion, Id: id}, nil
}

func (s *ToDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {

	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := s.Connect(ctx)
	if err != nil {
		return nil, sql.ErrConnDone
	}
	//從資料庫拿出來後要關掉
	rows, err := c.QueryContext(ctx, "SELECT `ID`,`Title`,`Description`,`Reminder` FROM ToDo WHERE `ID`=?", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "查詢失敗:"+err.Error())
	}
	defer rows.Close()
	//如果id有多筆	表示有問題
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("ID='#{req.Id}'找不到"))
		}
	}
	var td v1.ToDo
	var reminder time.Time
	if err := rows.Scan(&td.Id, &td.Title, &td.Description, &reminder); err != nil {
		return nil, status.Error(codes.Unknown, "查詢數據失敗:"+err.Error())
	}
	td.Reminder, err = ptypes.TimestampProto(reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "reminder格式無效:"+err.Error())
	}
	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("查找到多筆ID='%d'", req.Id))
	}
	return &v1.ReadResponse{Api: apiVersion, ToDo: &td}, nil
}

func (s *ToDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := s.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder參數無效")
	}

	res, err := c.ExecContext(ctx, "UPDATE ToDo SET `Title`=?, `Description`=?, `Reminder`= ? WHERE `ID=?", req.ToDo.Title, req.ToDo.Description, reminder, req.ToDo.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "更新失敗"+err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "更新row時失敗"+err.Error())
	}
	if rows == 0 {
		msg := "ID=" + strconv.FormatInt(req.ToDo.Id, 10) + "找不到"
		return nil, status.Error(codes.NotFound, msg)
	}
	return &v1.UpdateResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

func (s *ToDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := s.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	res, err := c.ExecContext(ctx, "DELETE FROM ToDo WHERE `ID`=?", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "刪除失敗"+err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "刪除該行失敗"+err.Error())
	}
	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("未找到ID='%id'", req.Id))
	}
	return &v1.DeleteResponse{Api: req.Api, Deleted: rows}, nil
}

func (s *ToDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := s.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo")
	if err != nil {
		return nil, status.Error(codes.Unknown, "查詢失敗"+err.Error())
	}
	defer rows.Close()

	var reminder time.Time
	list := []*v1.ToDo{}
	for rows.Next() {
		td := new(v1.ToDo)
		if err := rows.Scan(&td.Id, &td.Title, &td.Description, &reminder); err != nil {
			return nil, status.Error(codes.Unknown, "查詢失敗"+err.Error())
		}
		r, err := ptypes.TimestampProto(reminder)
		td.Reminder = r
		if err != nil {
			list = append(list, td)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "遍歷row時失敗"+err.Error())
	}
	return &v1.ReadAllResponse{
		Api: apiVersion,
		Set: list,
	}, nil
}
