# Golang

- [Why Golang](#why-golang)
- [Install](#install)
- [Run go program](#run-go-program)
	- [go build to Run](#go-build-to-run)
	- [go run to Run](#go-run-to-run)
	- [go install to Run](#go-install-to-run)
- [Code Format Tools](#code-format-tools)
	- [golint](#golint)
	- [goimports](#goimports)
- [Go Basic](#go-basic)
	- [Type Declarations](#type-declarations)
	- [Functions](#functions)
	- [Interfaces](#interfaces)
	- [Inheritance](#inheritance)
	- [Concurrency](#concurrency)
		- [goroutine](#goroutine)
		- [channel](#channel)
- [GOPATH vs GOROOT vs GOBIN vs Go Module](#gopath-vs-goroot-vs-gobin-vs-go-module)
- [go mod - modular programming, Go packages, Go modules](#go-mod---modular-programming-go-packages-go-modules)
	- [Go Package](#go-package)
- [BDD Testing Framework for Go](#bdd-testing-framework-for-go)
- [Resources](#resources)

## Why Golang

1. 简单
2. 代码一致性规范，易读
3. 突出的性能

## [Install](docs/installation.md)

## Run go program

### go build to Run

```sh

go build main.go

./main


```

### go run to Run

```sh
# or directly use 'go run' command.
# 'go run' command is not recommended to compile and run large Go projects.
# It is just a convenient way to run simple Go programs.

go run hello.go 

```

### go install to Run

```sh
go install  

# go install: no install location for directory /home/naveen/Documents/learngo outside GOPATH  
# For more details see: 'go help gopath'

export GOBIN=~/go/bin/

go install

~/go/bin/golang-getstart

```


## Code Format Tools

In VSCode, **Ctrl+Shift+P** to Open "Go:Install/Update Tool" and install the golang format tools:

### golint
### goimports

- Add import
- Organize import 

## Go Basic

### Type Declarations

```go

package models

// Student is a student object
type Student struct {
	ID int
	FirstName string
	LastName string
	Age int
}

```

```C#

// NewStudent is a constructor function that will return a Student pointer
func NewStudent(firstName string, lastName string, age int) *Student {
	return &Student{
		FirstName : firstName,
		LastName : lastName,
		Age: age,
	}
}


```

### Functions

- 函数名首字母要大写

### Interfaces

Interfaces are a pure object of design they just define a set of behaviors (ie. Methods) without giving any implementation of those behaviors.

```go

type DataRepository interface {
	Add(s *Student) (error)
	Update(s *Student) (error)
	Get(id int) (*Student, error)
	Delete(id int) (error)
}

```

```C#
// c#
public interface IDataRepository<T>
{
    void Add(T s);
    void Update(T s);
    T Get(int id);
    void Delete(int id);
}
```

### Inheritance

```go
// go
type DataReaderWriter interface {
	io.Reader
	io.Writer
}

type DataWriterReaderImpl struct {
}
// implements io.Reader
func (dwr DataWriterReaderImpl) Read(p []byte) (n int, err error) {
	//
	return 0, nil
}

// implements io.Writer
func (dwr DataWriterReaderImpl) Write(p []byte) (n int, err error) {
	//
	return 0, nil
}

```

### Concurrency

#### goroutine

Go 运行时管理的轻量级线程。

Benefits of Goroutines:

- With Goroutines, concurrency is achieved in Go programming. It helps two or more independent functions to run together.
- Goroutines can be used to run background operations in a program.
- It communicates through private channels so the communication between them is safer.
- With goroutines, we can split one task into different segments to perform better.

#### channel

Channels in Go act as a medium for goroutines to communicate with each other.

In Go, we use the make() function to create a channel. For example,

```go

channelName := make(chan int)

```


## GOPATH vs GOROOT vs GOBIN vs Go Module


**GOPATH** environment variable specifies the location of your workspace.
If no GOPATH is set, it is assumed to be $HOME/go on Unix systems and %USERPROFILE%\go on Windows.

```sh

# set WorkSpace GOPATH
export GOPATH=$PWD
cd $GOPATH && mkdir /bin && mkdir /pkg
export PATH=$GOPATH/bin:$PATH





# Check the GOPATH default setting
go env

# set terminal GOPATH (for zsh terminal, .zshrc)

export GOPATH="/Users/lning/go"
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

```

**GOROOT** is a variable that defines where your Go SDK is located. You do not need to change this variable, unless you plan to use different Go versions. default /usr/local/go


## go mod - modular programming, Go packages, Go modules

modular programming, go module is a collection of packages stored in a file tree with a go.mod file.

```sh

go mod init github.com/liuning0820/golang-basic
# run the module
go run github.com/liuning0820/golang-basic

# tidy: add missing and remove unused modules
go mod tidy

```

### Go Package

A package is a collection of source files in the same directory that are compiled together.
Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.
## BDD Testing Framework for Go 

https://github.com/onsi/ginkgo

## Resources

- [ ] [1] <https://go.dev/>
- [ ] <https://go.dev/doc/code>
- [ ] [go.dev/blog](https://go.dev/blog/pprof)
- [x] [pkg.go.dev - Discover Packages](https://pkg.go.dev/)
- [x] <https://www.practical-go-lessons.com>
- [x] <https://gobyexample.com>
- [ ] <https://golang.google.cn/>
- [ ] <https://github.com/golang-standards/project-layout>
- [ ] <https://github.com/golang/go/wiki/>
- [ ] <https://golangdocs.com/>

- [ ] https://tour.go-zh.org/methods/1

- [ ] https://go-zh.org/pkg/net/http/

- [ ] <https://studygolang.gitbook.io/learn-go-with-tests/>

- [ ] https://awesome-go.com/
- [](https://play-with-go.dev/go-fundamentals_go115_en/)
- [ ] [go and C# comparison]<https://devblogs.microsoft.com/premier-developer/go-and-csharp-comparison/>
- [ ] <https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go>
- [ ] <https://studygolang.gitbook.io/learn-go-with-tests>
