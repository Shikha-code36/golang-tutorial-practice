# Golang Tutorial

## Introduction to Go

Go, also known as Golang, is an open-source programming language developed by Google. It's known for its simplicity, strong static typing, and efficiency. Some of the main features of Go include:

- **Concurrency**: Go has built-in support for concurrent programming with goroutines and channels.
- **Garbage Collection**: Go manages memory automatically, freeing the developer from manual memory management.
- **Fast Compilation**: Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection.

## Package vs Module

In Go, a **package** is a collection of source files in the same directory that are compiled together. Functions, types, variables, or constants defined in one source file are visible to all other source files within the same package.

A **module** is a collection of related Go packages that are released together. A Go module is defined by a `go.mod` file at the root of the module directory. This file defines the module path and its dependency requirements.

## Initializing a Go Module

To initialize a new module, use the `go mod init` command followed by the module path. For example:

```shell
go mod init github.com/Shikha-code36/golang-tutorial-practice
```

After running the command `go mod init github.com/Shikha-code36/golang-tutorial-practice`, a `go.mod` file will be created in the current directory. The module path will be set to `github.com/Shikha-code36/golang-tutorial-practice`.

## Building and Running the Go Program

Once the module is initialized, you can build and run your Go program. This involves a few steps:

1. **Building the Program**: This step compiles the Go source file `main.go` located in the `tutorial1` directory.

```shell
go build tutorial1/main.go
.\main.exe
```

2. **Running the Program**: After building, you can run the resulting executable. Alternatively, you can directly run the program without creating an executable.

```shell
go run tutorial1/main.go
```

### Topics covered:
- [x] [Basics](tutorial1)
- [x] [Constants Variables and Basic Data Types](tutorial2)
- [x] [Functions and Control Structures](tutorial3)
- [x] [Arrays, Slices, Maps and Loops](tutorial4)
- [x] [Strings, Runes, and UTF-8 Encoding](tutorial5)
- [x] [Structs and Interfaces](tutorial6)
- [x] [Pointers](tutorial7)
- [x] [Concurrency(Goroutines)](tutorial8)