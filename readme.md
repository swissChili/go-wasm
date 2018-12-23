# go-wasm
A Go build utility for WebAssembly projects

## Building
```sh
go build -o wasm *.go
```

## Usage
go-wasm reads from a `wasm.yml` file in your current directory. Here is an example `wasm.yml` file:
```yml
# directory to use as static dir for HTML
static: static
# Go files to build
source: "*.go"
# Target to compile to
target: wasm
# Directory to read raw, uncompiled CSS from
cssdir: sass
# compiler 
csscomp:
  # OUTPUT and INPUT are placeholders for the args to use
  sass: sass INPUT OUTPUT
  scss: sass INPUT OUTPUT
# Output dir for compiled, static HTML, CSS, JS and wasm
output: target
```