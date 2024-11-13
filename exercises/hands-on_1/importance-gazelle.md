Ah, I see the issue. The error occurs because we included `go_googleapis` in the MODULE.bazel file but it's not actually being used in our project. Let's fix the MODULE.bazel file by removing the unnecessary dependency.



Update your MODULE.bazel file to look like this:

```python
# MODULE.bazel
module(
    name = "modern_bazel_demo",
    version = "0.1",
)

bazel_dep(name = "rules_go", version = "0.42.0")
bazel_dep(name = "gazelle", version = "0.35.0")

go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")
use_repo(go_deps)  # Remove the specific go_googleapis reference
```

Now try running:
```bash
bazel mod tidy
bazel run //:gazelle
```

The key changes are:
1. Removed the specific `go_googleapis` from `use_repo`
2. Left `use_repo(go_deps)` without specific repositories
3. Added `bazel mod tidy` step to clean up dependencies

This approach lets Gazelle automatically manage the Go dependencies based on what's actually used in your go.mod file. When you actually need to add external dependencies later, you would:

1. First add them to your Go project using `go get`
2. Update your go.mod with `go mod tidy`
3. Run `bazel run //:gazelle -- update-repos -from_file=go.mod`
4. Run `bazel mod tidy`
5. Run `bazel run //:gazelle`

___

Great! Let's add a popular logging library called `zap` from Uber as an example. We'll modify our project to use structured logging.

1. First, let's add zap to our Go project:
```bash
go get go.uber.org/zap
go mod tidy
```

2. Now let's update our code to use zap:

```go
// cmd/hello/main.go
package main

import (
    "example.com/modern-bazel-demo/pkg/calculator"
    "go.uber.org/zap"
)

func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    result := calculator.Add(5, 3)
    logger.Info("Calculation performed",
        zap.Int("a", 5),
        zap.Int("b", 3),
        zap.Int("result", result),
    )
}
```

3. Update the calculator package to use logging:

```go
// pkg/calculator/calculator.go
package calculator

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
    logger, _ = zap.NewProduction()
}

func Add(a, b int) int {
    result := a + b
    logger.Debug("Add operation performed",
        zap.Int("a", a),
        zap.Int("b", b),
        zap.Int("result", result),
    )
    return result
}

func Subtract(a, b int) int {
    result := a - b
    logger.Debug("Subtract operation performed",
        zap.Int("a", a),
        zap.Int("b", b),
        zap.Int("result", result),
    )
    return result
}
```

4. Now, let's update our Bazel dependencies. Run these commands:
```bash
# Update Gazelle with new dependencies
bazel run //:gazelle -- update-repos -from_file=go.mod

# Run mod tidy to clean up dependencies
bazel mod tidy

# Update BUILD files
bazel run //:gazelle
```

5. Verify everything works:
```bash
# Build the project
bazel build //cmd/hello:hello

# Run the binary
bazel run //cmd/hello:hello
```

You should see structured JSON logs like this:
```json
{"level":"info","ts":"2023-11-12T12:34:56.789Z","caller":"hello/main.go:15","msg":"Calculation performed","a":5,"b":3,"result":8}
```

Let's add a more complex example using zap's custom logger configuration:

```go
// cmd/hello/main.go
package main

import (
    "example.com/modern-bazel-demo/pkg/calculator"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func newLogger() *zap.Logger {
    config := zap.Config{
        Encoding:         "json",
        Level:           zap.NewAtomicLevelAt(zap.DebugLevel),
        OutputPaths:     []string{"stdout"},
        ErrorOutputPaths: []string{"stderr"},
        EncoderConfig: zapcore.EncoderConfig{
            MessageKey:     "message",
            LevelKey:      "level",
            TimeKey:       "time",
            EncodeLevel:   zapcore.CapitalLevelEncoder,
            EncodeTime:    zapcore.ISO8601TimeEncoder,
            EncodeCaller:  zapcore.ShortCallerEncoder,
        },
    }
    
    logger, _ := config.Build()
    return logger
}

func main() {
    logger := newLogger()
    defer logger.Sync()

    logger.Info("Starting calculator application")
    
    // Perform multiple calculations
    numbers := [][2]int{{5, 3}, {10, 7}, {15, 8}}
    
    for _, pair := range numbers {
        a, b := pair[0], pair[1]
        
        addResult := calculator.Add(a, b)
        logger.Info("Addition performed",
            zap.Int("a", a),
            zap.Int("b", b),
            zap.Int("result", addResult),
        )
        
        subResult := calculator.Subtract(a, b)
        logger.Info("Subtraction performed",
            zap.Int("a", a),
            zap.Int("b", b),
            zap.Int("result", subResult),
        )
    }
    
    logger.Info("Calculator application finished")
}
```

Now when you run:
```bash
bazel run //cmd/hello:hello
```

You'll see nicely formatted structured logs for all operations.

Some key points about what we just did:

1. **Dependency Management**:
   - Gazelle automatically handled the zap dependency in Bazel
   - The BUILD files were automatically updated with the correct dependencies

2. **Structured Logging**:
   - We're now using proper structured logging instead of fmt.Printf
   - All operations are tracked with relevant metadata
   - Logs are in JSON format for easy parsing

3. **Configuration**:
   - We added custom logger configuration
   - Debug level is enabled for more verbose output
   - Time format is ISO8601 for standard compliance

Would you like to:
1. Add more complex logging scenarios?
2. Add more dependencies to see how Gazelle handles them?
3. Explore other aspects of dependency management?

Ah, for the transition period between WORKSPACE and MODULE.bazel, we still need an empty WORKSPACE file. This is because some tools (including Gazelle) still expect it to exist. Let's fix this:

1. First, create an empty WORKSPACE file:
```bash
# Create empty WORKSPACE file
touch WORKSPACE
```

2. Then run the commands again:
```bash
# Update Gazelle with new dependencies
bazel run //:gazelle -- update-repos -from_file=go.mod

# Run mod tidy to clean up dependencies
bazel mod tidy

# Update BUILD files
bazel run //:gazelle
```

This is a common gotcha during the transition to Bzlmod. Even though we're using MODULE.bazel for our dependency management, some tools in the Bazel ecosystem still expect a WORKSPACE file to exist, even if it's empty.

The reason for this is:
- Some tools and rules haven't fully migrated to Bzlmod yet
- The WORKSPACE file is still used as a marker to identify the root of a Bazel workspace
- It provides backwards compatibility during the transition period

After creating the empty WORKSPACE file, you should be able to proceed with the dependency updates. Would you like to try running the commands again?

> I got an error with `zap` dependencies

The error occurs because we're missing the zapcore dependency in our BUILD.bazel file. Let's fix this by updating the dependencies.

1. First, let's modify the BUILD.bazel file in cmd/hello:

```python
# cmd/hello/BUILD.bazel
load("@rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "hello",
    srcs = ["main.go"],
    deps = [
        "//pkg/calculator",
        "@org_uber_go_zap//:zap",
        "@org_uber_go_zap//zapcore",  # Add zapcore dependency
    ],
    visibility = ["//visibility:public"],
)
```

2. Then update the MODULE.bazel file to ensure we have all needed dependencies:

```python
# MODULE.bazel
module(
    name = "modern_bazel_demo",
    version = "0.1",
)

bazel_dep(name = "rules_go", version = "0.42.0")
bazel_dep(name = "gazelle", version = "0.35.0")

go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")
use_repo(
    go_deps,
    "org_uber_go_zap",
    "org_uber_go_multierr",  # zap dependency
    "org_uber_go_atomic",    # zap dependency
)
```

3. Run the following commands to update everything:

```bash
# Update dependencies
bazel mod tidy

# Run gazelle to update BUILD files
bazel run //:gazelle

# Try building again
bazel build //cmd/hello:hello
```

The error occurred because zapcore is a separate package within zap that needs to be explicitly declared as a dependency. After making these changes, the build should work.

Also, let's verify that your go.mod has all the necessary dependencies:

```go
// go.mod should include something like:
require (
    go.uber.org/zap v1.26.0
    go.uber.org/atomic v1.11.0 // indirect
    go.uber.org/multierr v1.11.0 // indirect
)
```

If you're still having issues, you can try cleaning the Bazel cache and rebuilding:

```bash
bazel clean --expunge
bazel build //cmd/hello:hello
```
