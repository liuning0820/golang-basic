# Golang

TOC:

- [Install-In-Linux](#install-in-linux)
  - [Set modules mirror proxy](#set-modules-mirror-proxy)
- [Run go program](#run-go-program)
- [Edit and Format](#code-format-tools)

## Install in Linux

Download the Go from <https://gomirrors.org/>

```sh

curl -O https://gomirrors.org/dl/go/go1.14.4.linux-amd64.tar.gz


#Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go. For example:

sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz

# Create workspace $(GOPATH) and add it's /bin into $PATH
mkdir -p $HOME/go/{bin,src}

nano ~/.profile
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
source ~/.profile

# Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:

nano ~/.profile
# and put the line below at the bottom
export PATH=$PATH:/usr/local/go/bin
# Note: changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

# or simply run below line
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile && source ~/.profile

go version

```

### Set modules mirror proxy

The [goproxy.io](https://goproxy.io/) is one of the world's earliest Go modules mirror proxy services.

```sh
# GOPROXY="https://proxy.golang.org,direct"
# GOPROXY="https://mirrors.aliyun.com/goproxy"

go env -w GO111MODULE=on
go env -w GOPROXY="https://goproxy.io,direct"

```

## Install on Mac
```shell

$ go version

```


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




## Working with Primitive Data Type

## Working with Collections

## Functions

- 函数名首字母要大写

## Controlling Program Flow

## GOPATH vs GOROOT vs GOBIN vs Go Module


**GOPATH** environment variable specifies the location of your workspace.
If no GOPATH is set, it is assumed to be $HOME/go on Unix systems and %USERPROFILE%\go on Windows.

```sh

# set WorkSpace GOPATH
export GOPATH=$PWD
cd $GOPATH && mkdir /bin && mkdir /pkg
export PATH=$GOPATH/bin:$PATH

# Check the GOPATH setting
go env

```

**GOROOT** is a variable that defines where your Go SDK is located. You do not need to change this variable, unless you plan to use different Go versions. default /usr/local/go


## modular programming, Go packages, Go modules

modular programming, go module is a collection of packages stored in a file tree with a go.mod file.
### go.mod

```sh

go mod init github.com/liuning0820/golang-basic
# run the module
go run github.com/liuning0820/golang-basic

# tidy: add missing and remove unused modules
go mod tidy

```

### Package initialization

Package initialization takes place into **init** function.





## BDD Testing Framework for Go 

https://github.com/onsi/ginkgo

## Resources

- [x] <https://www.practical-go-lessons.com>
- [] <https://golang.google.cn/>

- [ ] <https://gobyexample.com>
- [ ] <https://golangdocs.com/>

- [ ] https://tour.go-zh.org/methods/1

- [ ] https://go-zh.org/pkg/net/http/

- [ ] <https://studygolang.gitbook.io/learn-go-with-tests/>

- [ ] <https://blog.golang.org/pprof>
- [ ] <https://blog.golang.org/slices-intro>

- [ ] <https://go.dev/>

- [](https://pkg.go.dev/)

- [ Discover Packages ](https://pkg.go.dev/)

- [](https://play-with-go.dev/go-fundamentals_go115_en/)

- [ ] <https://github.com/golang/go/wiki/>
