package main

import (
	fp "path/filepath"
	"os"
	"strings"
)

func compileCSS(sourceDir string, output string, compiler CssCompiler) error {
	return fp.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f := strings.Split(path, "/")
			file := f[len(f)-1]
			e := strings.Split(file, ".")
			extension := e[len(e)-1]
			command := ""
			comp := ""

			if extension == "sass" {
				comp = compiler.Sass
			} else if extension == "scss" {
				comp = compiler.Scss
			} else if extension == "less" {
				comp = compiler.Less
			}

			if comp != "" {
				file = file[:len(file)-5]
				command = strings.Replace(comp, "INPUT", path, -1)
				command = strings.Replace(command, "OUTPUT", output + "/" + file + ".css", -1)
			}

			status("Building Styles", command)
			err = runCommand(command)
			return err
		}
		return nil
	})
}