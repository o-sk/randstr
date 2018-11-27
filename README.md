# randstr
`randstr` is a command to copy a random character string to clipboard.

[![CircleCI](https://circleci.com/gh/o-sk/randstr/tree/master.svg?style=svg)](https://circleci.com/gh/o-sk/randstr/tree/master)

## Installation
Use `go get` to install this package:

```bash
$ go get github.com/o-sk/randstr
```

## Usage
```bash
$ randstr help
NAME:
   randstr - Copy random strings to clipboard

USAGE:
   randstr [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     alphabet, a  alphabet only
     number, n    number only
     help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --length value, -l value  Length of random string (default: 20)
   -o                        Output stdout
   --help, -h                show help
   --version, -v             print the version
```