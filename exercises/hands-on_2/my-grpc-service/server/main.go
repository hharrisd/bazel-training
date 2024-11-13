package main

import (
    "context"
    "log"
    "net"

    "example.com/my-grpc-service/proto/greeter"
    "google.golang.org/grpc"
)

type server struct {
    greeter.UnimplementedGreeterServiceServer
}

func (s *server) SayHello(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloResponse, error) {
    return &greeter.HelloResponse{Message: "Hello, " + in.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    greeter.RegisterGreeterServiceServer(s, &server{})

    log.Println("Server listening on :8080")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}