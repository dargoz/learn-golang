# Go Modules vs Packages

## What is a Package?
- A package is a collection of Go source files in the same directory that are compiled together.
- Each file starts with a `package` declaration (e.g., `package main`, `package math`).
- Packages help organize code and enable code reuse.
- You import packages in your code using the `import` statement.

**Example:**
```go
package math

func Add(a, b int) int {
    return a + b
}
```

## What is a Module?
- A module is a collection of related Go packages, managed as a single unit.
- A module is defined by a `go.mod` file at its root, which tracks the module’s path and its dependencies.
- Modules make it easy to manage dependencies and versioning for your project.

**Example:**
A project folder with a `go.mod` file and several packages:
```
myproject/        ← module root (contains go.mod)
    go.mod
    main.go       ← package main
    utils/
        math.go   ← package utils
```

## Key Differences
- **Package:** Smallest unit of code organization and reuse. Each directory with Go files and a `package` statement is a package.
- **Module:** A collection of packages, managed together for dependency and version control. Each project typically has one module.

## Summary Table
| Aspect    | Package                | Module                        |
|-----------|------------------------|-------------------------------|
| Scope     | Single directory       | Multiple packages/directories |
| Defined by| package statement      | go.mod file                   |
| Purpose   | Code organization      | Dependency/version management |
