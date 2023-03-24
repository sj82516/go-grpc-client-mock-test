package __

import (
    "context"
    
    "google.golang.org/grpc"
)

type client struct {
    grpcClient ExampleServiceClient
}

func NewClient(grpcHost string) *client {
    conn, err := grpc.Dial(grpcHost, grpc.WithInsecure())
    if err != nil {
        panic(err)
    }
    
    return &client{
        grpcClient: NewExampleServiceClient(conn),
    }
}

func (c *client) GetExample(ctx context.Context, name string) (*ExampleResponse, error) {
    return c.grpcClient.ExampleMethod(ctx, &ExampleRequest{
        Name: name,
    })
}
