package main

import (
	"syscall/js"
)

func main() {
	js.Global().Get("document").Call("getElementById", "main").Set("innerHTML", `
	<h1>Hello, World</h1>
	<h1>This hurts my eyes</h1>`)
	println("Hello There")
}