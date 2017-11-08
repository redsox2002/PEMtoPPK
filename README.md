# PEMtoPPK
A Go CLI tool to convert PEM files to PPK files for PuTTY using the cli package provided by [urfave/cli](https://github.com/urfave/cli)

When you create a key pair in AWS and download it to SSH into a server, it will give you a `.pem` file. Now this is perfectly fine for NIX based systems, but for Windows this isn't as friendly. If you are a Windows user and use [Putty](http://www.putty.org/), you know that you need a PPK file. This tool allows you to convert a PEM file to a PPK file quickly and easily instead of using PuttyGen (we all know CLI tools are quicker) :grin:

### Prerequisites:

- Install puttygen via Homebrew on Mac `brew install putty`, on Windows PC's PuTTYgen is installed when you install PuTTY
- Download the Go binary for your system here: https://golang.org/dl/.
- Follow the [installation instructions](https://golang.org/doc/install) for your system.
- Install `cli` by running `go get github.com/urfave/cli`

### Usage:

1. Build the main package by running `go build main.go`
2. Run `./main.go convert --pem /path/to/pem --ppk /path/to/save/ppk`
  - `/path/to/pem` is where you have saved your `.pem` key
  - `/path/to/save/ppk` is where you want to save your `.ppk` key
