# Golang

## Install in Linux


```sh

#Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go. For example:

tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz

export PATH=$PATH:/usr/local/go/bin

# Add Environments in your profile （optional）
# nano ~/.profile
# and put the line below at the bottom
PATH=$PATH:/usr/local/go/bin

# or simply run below line
# echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile && source ~/.profile


go version

```

## Set modules mirror proxy

The [goproxy.io](https://goproxy.io/) is one of the world's earliest Go modules mirror proxy services.

```sh
# GOPROXY="https://proxy.golang.org,direct"
# GOPROXY="https://mirrors.aliyun.com/goproxy"

go env -w GO111MODULE=on
go env -w GOPROXY="https://goproxy.io,direct"

```
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



