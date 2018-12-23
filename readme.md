# go-wasm
A Go build utility for WebAssembly projects

## Building
```sh
go build -o wasm *.go
```

## Usage
> **IMPORTANT**: Please add a script tag with source `index.js` in your `index.html` file! This is
> used to load your web assembly

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
  less: lessc INPUT -o OUTPUT
# Output dir for compiled, static HTML, CSS, JS and wasm
output: target
```
It will empty the output directory, compile your Go files to wasm, copy your static files over
(anything in the static dir will be copied, use this for images and such), compile your sass /
scss / less, and finally generate a JavaScript file that mounts your wasm program to the page. 

## Example
The `example` directory contains a small program that can be compiled for web assembly. The 
`wasm.yml` in the root contains the instructions to compile it. Feel free to look through the
`static/index.html` and `sass/main.sass` files, and the outputed files to see what changes
compiling makes.

## Server
Did I mention this has a nifty little server to serve up your fresh-off-the-compiler wasm project?
Well it does! No need to configure anything, it will automatically serve from the directory you
built to, and defaults to port `8080`. Use the `-port` flag to use a different port
> eg: `wasm -port 9090`