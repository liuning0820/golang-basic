# Golang

TOC:

- [Install](#install-in-linux)
  - [Set modules mirror proxy](#set-modules-mirror-proxy)
- [Run go program](#run-go-program)
- [Edit and Format](#code-format-tools)

## Install in Linux

Download the Go from <https://gomirrors.org/>

```sh

#Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go. For example:

tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz

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

## Run go program

### go build to Run

```sh

go build hello.go

./hello


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

## go mod

```sh

go mod init github.com/liuning0820/golang-basic
# run the module
go run github.com/liuning0820/golang-basic

```


## Working with Primitive Data Type

## Working with Collections

## Functions

- 函数名首字母要大写

## Controlling Program Flow

## GOPATH vs GOROOT vs GOBIN

导入路径是一个字符串，用户可以用它来导入包。它指定了包源码所在的目录（相对于 $GOROOT/src/pkg 或 `$GOPATH/src`)。

**GOPATH** environment variable specifies the location of your workspace.
If no GOPATH is set, it is assumed to be $HOME/go on Unix systems and %USERPROFILE%\go on Windows.

**GOROOT** is a variable that defines where your Go SDK is located. You do not need to change this variable, unless you plan to use different Go versions.




## Resources

- [ ] <https://gobyexample.com>

- [ ] https://tour.go-zh.org/methods/1

- [ ] https://go-zh.org/pkg/net/http/

- [ ] <https://studygolang.gitbook.io/learn-go-with-tests/>

- [ ] <https://blog.golang.org/pprof>
- [ ] <https://blog.golang.org/slices-intro>

- [ ] <https://go.dev/>

- [ ] <https://pkg.go.dev/>

- [ ] <https://github.com/golang/go/wiki/>
