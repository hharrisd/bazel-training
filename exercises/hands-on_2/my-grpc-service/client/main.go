package main

import (
    "context"
    "log"

    "example.com/my-grpc-service/proto/greeter"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    c := greeter.NewGreeterServiceClient(conn)

    resp, err := c.SayHello(context.Background(), &greeter.HelloRequest{Name: "Alice"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", resp.GetMessage())
}