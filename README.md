# Golang

## Install in Linux


```sh

#Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go. For example:

tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz

export PATH=$PATH:/usr/local/go/bin

go version

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

