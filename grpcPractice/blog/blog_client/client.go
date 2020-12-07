package main

import (
	"context"
	"fmt"
	"grpcPractice/blog/blogpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("BLog Client")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50053", opts)
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}
	defer cc.Close()
	c := blogpb.NewBlogServiceClient(cc)
	createBlogRes := CreateBlogRes(c)
	blogID := createBlogRes.GetBlog().GetId()

	ReadBlog(c, blogID)
	UpdateBlog(c, blogID)
	Delete(c, blogID)
	ListBlogs(c)
}

func CreateBlogRes(c blogpb.BlogServiceClient) *blogpb.CreateBlogResponse {
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Howard",
		Title:    "My First blog",
		Content:  "Contents",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}
	fmt.Printf("Blog has been created: %v\n", createBlogRes)
	return createBlogRes
}

func ReadBlog(c blogpb.BlogServiceClient, blogID string) {
	fmt.Println("Reading the blog")
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: ""})
	if err2 != nil {
		fmt.Printf("Error happened while reading empty id: %v \n", err2)
	}
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readBlogErr)
	}
	fmt.Printf("Blog was read: %v \n", readBlogRes)
}

func UpdateBlog(c blogpb.BlogServiceClient, blogID string) {
	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Changed Author",
		Title:    "edited",
		Content:  "update content",
	}
	res, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if err != nil {
		fmt.Printf("Error happened while updating: %v \n", err)
	}
	fmt.Printf("Blog was updated: %v\n", res)
}

func Delete(c blogpb.BlogServiceClient, blogID string) {
	res, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})
	if err != nil {
		fmt.Printf("Error happened while deleting %v\n", err)
	}
	fmt.Printf("Blog was deleted: %v \n", res)
}

func ListBlogs(c blogpb.BlogServiceClient) {
	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequuest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Stream crushed: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
