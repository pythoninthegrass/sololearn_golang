# SoloLearn - Go

## Setup
```bash
# install homebrew (macOS)
bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# install asdf
brew install asdf

# install go
asdf plugin-add golang https://github.com/kennyp/asdf-golang
asdf install golang latest
asdf global golang latest

# add asdf shims to ~/.bashrc
# append $GOPATH to existing $PATH (default macOS path listed)
export GOPATH=$HOME/asdf/shims
export PATH="$GOPATH:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:$PATH"
```

## Usage
```bash
# create directory and move to it
mkdir -p
cd hello/

# start module and track code's dependencies
go mod init example/hello

# edit hello.go

# run program in working directory
go run .
```

## TODO
* Add VSCode usage (plugins, debugging)

## Further reading
[Learn Go | Sololearn](https://www.sololearn.com/learning/1164)

[The Missing Package Manager for macOS (or Linux) — Homebrew](https://brew.sh/)

[Introduction | asdf](https://asdf-vm.com/guide/introduction.html)

[Tutorial: Get started with Go - The Go Programming Language](https://golang.org/doc/tutorial/getting-started)

[Tutorial: Create a Go module - The Go Programming Language](https://golang.org/doc/tutorial/create-module)

[pkg.go.dev](https://pkg.go.dev/)
