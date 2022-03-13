# Learning Golang

Golang Beginner
-----
### go hompage
https://go.dev/doc/install

```
go version
go version go1.17.8 darwin/arm64
```

GOROOT check
```
which go
/usr/local/go/bin/go
```
or
```
go env GOROOT
/usr/local/go
```

mkdir go
```
mkdir -p $HOME/go/{bin,src,pkg}
```

~/.zshrc write
```
export GOPATH=$HOME/go
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
```
source ~/.zshrc

-----
### mac brew

link : https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5

```
brew update&& brew install golang

go version
go version go1.17.8 darwin/arm64
```

GOROOT check
```
which go
/opt/homebrew/bin/go
```
or
```
go env GOROOT
/opt/homebrew/Cellar/go/1.17.8/libexec
```

mkdir go
```
mkdir -p $HOME/go/{bin,src,pkg}
```

~/.zshrc write
```
export GOPATH=$HOME/go
# export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
```
source ~/.zshrc

-----

### gvm install (like=nvm)

install
```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

~/.zshrc write
```
[[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"
```
source ~/.zshrc

default
```
gvm install go1.4 -B
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.5
```

mac m1
```
gvm install go1.10 -B
export GOROOT_BOOTSTRAP=$GOROOT
gvm use go1.10
gvm install go1.10.8
gvm use go1.10.8
```

-----
### gvm apply error issue

link : https://dc7303.github.io/go/2019/11/29/goGvmInstallVersion/

-----