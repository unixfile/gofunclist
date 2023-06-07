# gofunclist

`gofunclist` is a command-line tool that lists all functions defined in a Go package, including their full function signatures.

## Installation

First, clone this repository:

```sh
git clone https://github.com/yourusername/gofunclist.git
cd gofunclist
```

Next, build the tool:

```sh
go build -o gofunclist main.go
```

Then, place the tool in your $PATH.

## Usage

Run the `gofunclist` command followed by the path of the Go package you want to inspect:

```sh
gofunclist $GOROOT/src/fmt
```

The tool will print a list of all function signatures defined in the package, in the following format:

```sh
func ((receiver type)) FunctionName(arg type) (return type)
```

## License
[MIT](license.md)

## Thanks to ChatGPT
This tool was created with the help of ChatGPT. The creation process is saved [here](chat.md).
