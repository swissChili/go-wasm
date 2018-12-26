package main

import (
	"fmt"
)

func compileGo(output string, sources string) error {
	buildCmd := fmt.Sprintf("GOOS=js GOARCH=wasm go build -o %s/main.wasm %s", output, sources)
	status("Building Go", sources)
	return runCommand(buildCmd)
}