package client

import (
    "context"
    "fmt"
    pb "mock-grpc/proto"
    "net"
    "testing"
    
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
    "google.golang.org/grpc"
)

type TestClient struct {
    suite.Suite
    grpcServer *grpc.Server
    mock       *pb.MockExampleServiceServer
    client     Client
}

func TestClientSuite(t *testing.T) {
    suite.Run(t, new(TestClient))
}

func (s *TestClient) SetupSuite() {}

func (s *TestClient) TearDownSuite() {}

const grpcHost = "localhost:50052"

func (s *TestClient) SetupTest() {
    s.grpcServer = grpc.NewServer()
    s.mock = pb.NewMockExampleServiceServer(s.T())
    pb.RegisterExampleServiceServer(s.grpcServer, s.mock)
    
    lis, err := net.Listen("tcp", grpcHost)
    if err != nil {
        fmt.Println("failed to listen: ", err)
    }
    
    go func() {
        if err := s.grpcServer.Serve(lis); err != nil {
            fmt.Println("failed to serve: ", err)
        }
    }()
    
    s.client = *NewClient(grpcHost)
}

func (s *TestClient) TearDownTest() {
    s.mock.AssertExpectations(s.T())
    s.grpcServer.Stop()
}

func (s *TestClient) TestGetExample() {
    message := "Hello World"
    s.mock.On("ExampleMethod", mock.Anything, mock.Anything).Return(&pb.ExampleResponse{
        Message: message,
    }, nil)
    
    res, err := s.client.GetExample(context.Background(), "test")
    if err != nil {
        fmt.Println("failed to get example: ", err)
    }
    
    s.Equal(message, res.Message)
}
