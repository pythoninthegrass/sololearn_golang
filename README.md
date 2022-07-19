# SoloLearn - Go

**Table of Contents**
* [SoloLearn - Go](#sololearn---go)
  * [Setup](#setup)
  * [Usage](#usage)
  * [TODO](#todo)
  * [Further reading](#further-reading)

## Setup
```bash
# install homebrew (macOS)
bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# install asdf
brew install asdf

# install go
asdf plugin-add golang https://github.com/kennyp/asdf-golang
asdf install golang 1.18.4
asdf local golang 1.18.4

# add asdf shims to ~/.bashrc
# append $GOPATH to existing $PATH (default macOS path listed)
export GOPATH=$(go env GOPATH)
export GOROOT=$(go env GOROOT)
export GOBIN=$(go env GOBIN)
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOBIN
export PATH="/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:$PATH"
```

* VSCode
  From @ko31: > Open the Extensions Marketplace (Cmd+Shift+X), search Go and install it. > > [O]pen the Command Palette (Cmd+Shift+P) and run the Go: Install/Update Tools command. > > Install all the Go extensions listed there.
* Known good config

```json
{
    // OTHER CONFIG ITEMS ^^
    // GOLANG-SPECIFIC
    "go.inferGopath": false,
    "go.buildOnSave": "workspace",
    "go.lintOnSave": "package",
    "go.vetOnSave": "package",
    "go.buildTags": "",
    "go.buildFlags": [],
    "go.lintFlags": [],
    "go.vetFlags": [],
    "go.coverOnSave": false,
    "go.useCodeSnippetsOnFunctionSuggest": false,
    "go.formatTool": "gofmt",
    "go.gocodeAutoBuild": false,
    "go.useLanguageServer": true,
    "go.alternateTools": {
        "go-langserver": "gopls"
    },
    "[go]": {
        "editor.codeActionsOnSave": {
            "source.organizeImports": false
        },
        "editor.formatOnSave": true,
    },
}
```

## Usage

```bash
# start module and track code's dependencies in tld
go mod init git_username/repo_name

# create directory and move to it
mkdir -p hello && cd $_

# edit hello.go

# run program in working directory
go run .

# install dependency
go get rsc.io/quote

# cleanup imports
go mod tidy
```

## TODO

* TODO

## Further reading

[Learn Go | Sololearn](https://www.sololearn.com/learning/1164)

[The Missing Package Manager for macOS (or Linux) — Homebrew](https://brew.sh/)

[Introduction | asdf](https://asdf-vm.com/guide/introduction.html)

[asdf to manage multiple Golang on Mac – ookangzheng](https://www.ookangzheng.com/asdf-to-manage-multiple-golang-on-mac/)

[Tutorial: Get started with Go - The Go Programming Language](https://golang.org/doc/tutorial/getting-started)

[Tutorial: Create a Go module - The Go Programming Language](https://golang.org/doc/tutorial/create-module)

[pkg.go.dev](https://pkg.go.dev/)

[How to setup Golang with VSCode - DEV Community](https://dev.to/ko31/how-to-setup-golang-with-vscode-1i4i)

[Quick Start - GoReleaser](https://goreleaser.com/quick-start/)
