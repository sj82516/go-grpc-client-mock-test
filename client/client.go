package client

import (
    "context"
    "fmt"
    
    pb "mock-grpc/proto"
    
    "google.golang.org/grpc"
)

type Client struct {
    grpcClient pb.ExampleServiceClient
}

func NewClient(grpcHost string) *Client {
    conn, err := grpc.Dial(grpcHost, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        fmt.Println("did not connect: ", err)
    }
    
    return &Client{
        grpcClient: pb.NewExampleServiceClient(conn),
    }
}

func (c *Client) GetExample(ctx context.Context, name string) (*pb.ExampleResponse, error) {
    return c.grpcClient.ExampleMethod(ctx, &pb.ExampleRequest{
        Name: name,
    })
}
