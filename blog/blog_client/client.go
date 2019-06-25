package main

import (
	"context"
	"fmt"
	"grpc-go-course/blog/blogpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	blog := &blogpb.Blog{
		AuthorId: "Jon",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error %v", err)
	}
	fmt.Printf("Blog has been created: %v\n", createBlogRes)

	// Read blog

	fmt.Println("Reading the blog")
	_, readErr := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "fsafasfsafsa"})
	if readErr != nil {
		fmt.Printf("Error happened while reading: %v\n", readErr)
	}

	readBlogResp, readBlogErr := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: createBlogRes.GetBlog().GetId()})
	if readBlogErr != nil {
		fmt.Printf("Error happened while reading: %v\n", readBlogErr)
	}

	fmt.Printf("Blog was read: %v\n", readBlogResp)

	// Update blog
	updatedBlog := &blogpb.Blog{
		Id:       createBlogRes.GetBlog().GetId(),
		AuthorId: "Changed Author",
		Title:    "My first blog (edited)",
		Content:  "Content of the first blog, with new things",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: updatedBlog})

	if updateErr != nil {
		fmt.Printf("Error happened while updating %v\n", updateErr)
	}
	fmt.Printf("Blog was updated: %v\n", updateRes)

	// Delete blog

	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: createBlogRes.GetBlog().GetId()})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting %v\n", deleteErr)
	}

	fmt.Printf("Blog was deleted: %v\n", deleteRes)

	// List blogs

	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("Error while calling rpc %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// We've reached end of the stream
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		log.Println(msg.GetBlog())
	}
}
