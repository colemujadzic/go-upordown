# go-upordown

[![Travis CI](https://img.shields.io/travis/colemujadzic/go-tweets.svg?style=for-the-badge)](https://travis-ci.org/colemujadzic/go-tweets)

Check if a website is up or down with this command-line tool, written in Go using the Cobra CLI library

## Installation

#### Via Go

```console
$ go get github.com/colemujadzic/go-upordown
```

## Usage

```console
$ go-upordown -h
go-upordown is a command-line tool written in Go, using the Cobra CLI library.

It makes it easy for anyone to figure out if a specified website is available or not.

Usage:
  go-upordown [command]

Available Commands:
  check       Check the status of a website
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.go-upordown.yaml)
  -h, --help            help for go-upordown
  -t, --toggle          Help message for toggle

Use "go-upordown [command] --help" for more information about a command.
```
