# Basic Concepts (3 hours)

This section covers the fundamental building blocks of Bazel and how they work together. Here's a more detailed breakdown:

### Targets and Labels
- **Targets**: Bazel's basic unit of build. Can be a binary, library, test, or custom rule.
- **Labels**: Unique identifiers for targets, using the format `//package/path:target_name`.
  - Example: `//cmd/hello:hello` refers to the `hello` target in the `cmd/hello` package.
- Understanding how to reference targets across packages and workspaces.

### Dependencies
- Declaring dependencies between targets using the `deps` attribute.
- Direct vs. transitive dependencies.
- Dependency resolution and version conflicts.

### Visibility
- Controlling access to targets using the `visibility` attribute.
- Public vs. private targets.
- Visibility patterns like `//visibility:public` and `//package/...`.

### Build Rules
- Bazel's built-in rules for common build tasks (e.g., `go_binary`, `go_library`, `cc_library`).
- Defining custom build rules using Starlark.
- Understanding rule attributes and their purpose.

### Command Line Interface
- Bazel command structure (e.g., `bazel build`, `bazel test`, `bazel run`).
- Build configuration via `.bazelrc` files.
- Querying the build graph using `bazel query`.
