# Exercise Bazel + Protobuf

Sure, let's dive into an example of using Bazel with Protocol Buffers (protobuf), leveraging the Bzlmod module system and integrating external libraries.

In this example, we'll create a simple gRPC service that uses protobuf for message definitions and communication.

## Project Structure

Let's start by setting up the project structure:

```
my-grpc-service/
├── MODULE.bazel
├── proto/
│   ├── greeter.proto
│   └── BUILD.bazel
├── server/
│   ├── main.go
│   └── BUILD.bazel
└── client/
    ├── main.go
    └── BUILD.bazel
```

## Setting up the Bzlmod Module

In the `MODULE.bazel` file, we'll declare our project module and the required dependencies:

```python
module(
    name = "my-grpc-service",
    version = "0.1",
)

bazel_dep(name = "rules_go", version = "0.42.0")
bazel_dep(name = "gazelle", version = "0.35.0")
bazel_dep(name = "rules_proto", version = "5.3.0")

go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")

use_repo(
    go_deps,
    "com_github_golang_protobuf",
    "google_golang_org_grpc",
)
```

Here, we're using Bzlmod to manage the following dependencies:

1. `rules_go`: For building Go code.
2. `gazelle`: For generating Bazel BUILD files.
3. `rules_proto`: For building protobuf-based code.
4. `com_github_golang_protobuf`: The protobuf Go implementation.
5. `google_golang_org_grpc`: The gRPC Go implementation.

## Defining the Protobuf Service

In the `proto/greeter.proto` file, we'll define the protobuf message and service:

```protobuf
syntax = "proto3";

package greeter;

service GreeterService {
  rpc SayHello(HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}
```

## Building the Protobuf Code

Next, we'll create the `proto/BUILD.bazel` file to build the protobuf code:

```python
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@com_github_grpc_ecosystem_grpc_gateway//bazel:go_proto_library.bzl", "go_proto_library")

proto_library(
    name = "greeter_proto",
    srcs = ["greeter.proto"],
)

go_proto_library(
    name = "greeter_go_proto",
    importpath = "example.com/my-grpc-service/proto/greeter",
    proto = ":greeter_proto",
)
```

This defines a `proto_library` target for the protobuf definition and a `go_proto_library` target to generate the Go code.

## Implementing the Server

In the `server/main.go` file, we'll implement the gRPC server:

```go
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
```

And the corresponding `server/BUILD.bazel` file:

```python
load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "server_lib",
    srcs = ["main.go"],
    importpath = "example.com/my-grpc-service/server",
    deps = [
        "//proto/greeter:greeter_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@google_golang_org_grpc//:go_default_library",
    ],
)

go_binary(
    name = "server",
    embed = [":server_lib"],
    visibility = ["//visibility:public"],
)
```

## Implementing the Client

In the `client/main.go` file, we'll create a gRPC client:

```go
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
```

And the corresponding `client/BUILD.bazel` file:

```python
load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "client_lib",
    srcs = ["main.go"],
    importpath = "example.com/my-grpc-service/client",
    deps = [
        "//proto/greeter:greeter_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@google_golang_org_grpc//:go_default_library",
    ],
)

go_binary(
    name = "client",
    embed = [":client_lib"],
    visibility = ["//visibility:public"],
)
```

## Building and Running the Example

To build and run the example:

1. Generate the protobuf code:
```bash
bazel build //proto:greeter_go_proto
```

2. Build the server and client:
```bash
bazel build //server:server //client:client
```

3. Run the server:
```bash
bazel run //server:server
```

4. Run the client:
```bash
bazel run //client:client
```

This example demonstrates how to:

1. Use Bzlmod to manage dependencies, including external libraries like `com_github_golang_protobuf` and `google_golang_org_grpc`.
2. Define a protobuf service and generate the Go code using the `rules_proto` and `go_proto_library` rules.
3. Implement the gRPC server and client using the generated code.
4. Build and run the server and client using Bazel.

The key benefits of this approach are:

- **Dependency Management**: Bzlmod provides a modern, reliable, and consistent way to manage dependencies, including external protobuf-related libraries.
- **Protobuf Integration**: The `rules_proto` and `go_proto_library` rules make it easy to integrate protobuf definitions into your Bazel-based Go project.
- **Build Automation**: Bazel handles the entire build process, from protobuf code generation to building the server and client binaries.

This example can be easily extended to include more complex protobuf definitions, multiple services, and additional functionality. The principles demonstrated here can be applied to a wide range of Bazel-based projects that involve protobuf and gRPC.