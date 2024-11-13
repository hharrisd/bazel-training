Great! Let's create your first modern Bazel project with Go support. We'll do this step by step.



```bash
# First, create a new project directory and navigate into it
mkdir modern-bazel-demo
cd modern-bazel-demo

# Create basic project structure
mkdir -p cmd/hello
mkdir pkg/calculator

# Initialize a go.mod file
go mod init example.com/modern-bazel-demo

# Create MODULE.bazel file
cat > MODULE.bazel << 'EOF'
module(
    name = "modern_bazel_demo",
    version = "0.1",
)

bazel_dep(name = "rules_go", version = "0.42.0")
bazel_dep(name = "gazelle", version = "0.35.0")

go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")
use_repo(go_deps, "go_googleapis")
EOF

# Create root BUILD.bazel file
cat > BUILD.bazel << 'EOF'
load("@gazelle//:def.bzl", "gazelle")

# Gazelle configuration
gazelle(
    name = "gazelle",
    prefix = "example.com/modern-bazel-demo",
)
EOF

# Create main.go in cmd/hello
cat > cmd/hello/main.go << 'EOF'
package main

import (
    "fmt"
    "example.com/modern-bazel-demo/pkg/calculator"
)

func main() {
    result := calculator.Add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
}
EOF

# Create BUILD.bazel for cmd/hello
cat > cmd/hello/BUILD.bazel << 'EOF'
load("@rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "hello",
    srcs = ["main.go"],
    deps = ["//pkg/calculator"],
    visibility = ["//visibility:public"],
)
EOF

# Create calculator package
cat > pkg/calculator/calculator.go << 'EOF'
package calculator

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}
EOF

# Create calculator test
cat > pkg/calculator/calculator_test.go << 'EOF'
package calculator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(5, 3)
    if result != 8 {
        t.Errorf("Add(5, 3) = %d; want 8", result)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    if result != 2 {
        t.Errorf("Subtract(5, 3) = %d; want 2", result)
    }
}
EOF

# Create BUILD.bazel for calculator package
cat > pkg/calculator/BUILD.bazel << 'EOF'
load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "calculator",
    srcs = ["calculator.go"],
    importpath = "example.com/modern-bazel-demo/pkg/calculator",
    visibility = ["//visibility:public"],
)

go_test(
    name = "calculator_test",
    srcs = ["calculator_test.go"],
    embed = [":calculator"],
)
EOF

# Create .bazelrc with some useful defaults
cat > .bazelrc << 'EOF'
# Build in C++ compilation mode by default
build --compilation_mode=opt

# Use Go modules by default
build --enable_bzlmod

# Test output should be detailed by default
test --test_output=all

# Enable more error output
build --verbose_failures
EOF

```

Now that we have our project structure set up, let's go through how to use it:

1. **First, initialize Gazelle**:
```bash
bazel run //:gazelle
```

2. **Build the project**:
```bash
bazel build //cmd/hello:hello
```

3. **Run the binary**:
```bash
bazel run //cmd/hello:hello
```

4. **Run the tests**:
```bash
bazel test //pkg/calculator:calculator_test
```

Let's go through what we've created:

1. **Project Structure**:
   - `cmd/hello`: Contains the main application
   - `pkg/calculator`: A sample package with unit tests
   - Root level Bazel configuration files

2. **Key Files**:
   - `MODULE.bazel`: Modern dependency management
   - `BUILD.bazel` files: Build configurations
   - `.bazelrc`: Build settings
   - `go.mod`: Go module definition

3. **Dependencies**:
   - rules_go: For building Go code
   - gazelle: For generating BUILD files

Would you like me to explain any specific part in more detail? Or should we move on to adding more features like external dependencies or creating custom build rules?