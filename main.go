package main

import (
	"fmt"
	"io/ioutil"
	"github.com/go-yaml/yaml"
	//fp "path/filepath"
	"os/exec"
	. "github.com/logrusorgru/aurora"
	// am I really importing something just to copy a dir?
	// yep. 
	. "github.com/otiai10/copy"
	"strings"
	fp "path/filepath"
	"os"
)

func fileNotFound(err error) {
	if err != nil {
		fmt.Println(`Failed to read file 'wasm.yml'.
Make sure it exists and the current user has read permissions.`)
		panic(err)
	}
}

func status(short string, desc string) {
	if len(short) >= 16 {
		short = short[0:16]
	}
	fmt.Printf("%16s :: %s\n", Cyan(short), desc)
}

func main() {
	dat, err := ioutil.ReadFile("wasm.yml")
	fileNotFound(err)

	var wasm Config
	yaml.Unmarshal(dat, &wasm)

	// Move all of the static dir contents
	err = Copy(wasm.Static, wasm.Output)
	status("Copying Static", fmt.Sprintf("%s -> %s", wasm.Static, wasm.Output))
	if err != nil { panic(err) }

	// Compile the go files
	buildCmd := fmt.Sprintf("GOOS=js GOARCH=wasm go build -o %s/main.wasm %s", wasm.Output, wasm.Source)
	status("Building Go", wasm.Source)
	_, err = exec.Command("sh", "-c", buildCmd).Output()
	if err != nil { panic(err) }

	// Compile all the sass files
	err = fp.Walk(wasm.CssDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f := strings.Split(path, "/")
			file := f[len(f)-1]
			e := strings.Split(file, ".")
			extension := e[len(e)-1]
			command := ""
			comp := ""

			if extension == "sass" {
				comp = wasm.CssComp.Sass
			} else if extension == "scss" {
				comp = wasm.CssComp.Scss
			}
			
			if comp != "" {
				file = file[:len(file)-5]
				command = strings.Replace(comp, "INPUT", path, -1)
				command = strings.Replace(command, "OUTPUT", wasm.Output + "/" + file + ".css", -1)
			}

			status("Building Styles", command)
			_, err = exec.Command("sh", "-c", command).Output()
			return err
		}
		return nil
	})
	if err != nil { panic(err) }

}