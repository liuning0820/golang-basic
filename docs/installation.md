# GoLang Installation

- [GoLang Installation](#golang-installation)
  - [Install on Windows](#install-on-windows)
  - [Install on Linux](#install-on-linux)
    - [Set modules mirror proxy](#set-modules-mirror-proxy)
  - [Install on MacOS](#install-on-macos)

## Install on Windows



## Install on Linux

Download the Go from <https://gomirrors.org/>

```sh

curl -O https://gomirrors.org/dl/go/go1.18.8.linux-amd64.tar.gz

wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz


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

## Install on MacOS

```shell

~ go version

```