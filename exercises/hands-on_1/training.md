# Modern Bazel Training Program: Using Bzlmod
## Overview
This training program focuses on the modern Bazel approach using Bzlmod (Bazel Modules) for dependency management.

## Module 1: Modern Bazel Fundamentals
### 1.1 Setup and Installation (2 hours)
- Installing the latest Bazel version
- Understanding the new module system
- MODULE.bazel vs. legacy WORKSPACE
- Basic project structure

#### Hands-on Exercise 1:
```bash
# Create a new Bazel project with modules
mkdir modern-bazel-demo
cd modern-bazel-demo
touch MODULE.bazel
```

Create your first MODULE.bazel file:
```python
# MODULE.bazel
module(
    name = "my_project",
    version = "0.1",
)

bazel_dep(name = "rules_go", version = "0.39.1")
bazel_dep(name = "gazelle", version = "0.31.0")
```

Create your first BUILD.bazel file:
```python
# BUILD.bazel
load("@rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "hello",
    srcs = ["hello.go"],
)
```

### 1.2 Module System Basics (3 hours)
- Module declaration and versioning
- Direct and indirect dependencies
- Version resolution
- Override system
- Multiple modules in a project

#### Hands-on Exercise 2:
Working with multiple dependencies:
```python
# MODULE.bazel
module(
    name = "my_project",
    version = "0.1",
)

bazel_dep(name = "rules_go", version = "0.39.1")
bazel_dep(name = "gazelle", version = "0.31.0")

# Example of dependency override
single_version_override(
    module_name = "rules_go",
    version = "0.39.1",
)

# Using go_deps
go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")
use_repo(go_deps, "go_googleapis")
```

## Module 2: Advanced Module Management
### 2.1 Working with Go Modules (4 hours)
- Integration with go.mod
- Gazelle with Bzlmod
- Managing multiple Go modules
- Version compatibility

#### Hands-on Exercise 3:
Create a multi-module Go project:
```python
# MODULE.bazel
module(
    name = "my_project",
    version = "0.1",
)

bazel_dep(name = "rules_go", version = "0.39.1")

# Local module dependency
local_path_override(
    module_name = "my_lib",
    path = "./lib",
)
```

```go
// go.mod
module example.com/myproject

go 1.20

require (
    github.com/some/dependency v1.2.3
)
```

### 2.2 Module Extensions (3 hours)
- Creating custom module extensions
- Extension loading and usage
- Extension versioning
- Cross-module communication

## Module 3: Build Configuration in Modern Bazel
### 3.1 Modern Configuration Approaches (3 hours)
- .bazelrc with modules
- Build setting dependencies
- Platform configuration
- Toolchain registration

#### Hands-on Exercise 4:
```python
# MODULE.bazel with toolchains
module(
    name = "my_project",
    version = "0.1",
)

bazel_dep(name = "rules_go", version = "0.39.1")

register_toolchains(
    "//toolchains:go_toolchain",
)
```

### 3.2 Dependency Management Best Practices (3 hours)
- Version pinning
- Registry dependencies
- Git dependencies
- Archive dependencies
- Private registry setup

## Module 4: Testing and CI/CD with Modern Bazel
### 4.1 Modern Testing Setup (3 hours)
- Test module organization
- Test dependencies
- Coverage reporting
- Test caching

#### Hands-on Exercise 5:
```python
# tests/MODULE.bazel
module(
    name = "my_project_tests",
    version = "0.1",
)

bazel_dep(name = "rules_go", version = "0.39.1")
bazel_dep(name = "my_project", version = "0.1")
```

### 4.2 Modern CI/CD Integration (3 hours)
- GitHub Actions with Bzlmod
- Module caching in CI
- Remote execution setup
- Registry integration

## Module 5: Enterprise Features
### 5.1 Private Registry (4 hours)
- Setting up private module registry
- Authentication and authorization
- Version management
- Mirror configuration

### 5.2 Large-Scale Development (4 hours)
- Multiple teams coordination
- Version conflict resolution
- Module boundaries
- Performance optimization

## Final Project
Build a microservices application using:
- Multiple Bazel modules
- Custom module extensions
- Private registry
- CI/CD pipeline with module caching

## Additional Resources
- Bzlmod documentation
- Migration guides from WORKSPACE
- Registry documentation
- Community modules and extensions

## Time Allocation
- Total training time: ~32 hours
- Recommended pace: 2-3 weeks
- Daily hands-on practice: 2-3 hours