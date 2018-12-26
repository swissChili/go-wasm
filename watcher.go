package main

import (
	"fmt"
	"os"
)

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, info os.FileInfo, err error) error {

	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if info.Mode().IsDir() {
		return watcher.Add(path)
	}

	return err
}


func watchAndRecompileCSS(sourceDir string, output string, compiler CssCompiler) {
	for {
		select {
		case <-watcher.Events:
			compileCSS(sourceDir, output, compiler)
		case e := <-watcher.Errors:
			fmt.Println("ERR", e)
		}
	}
}

func watchAndRecompileGo(output string, sources string) {
	for {
		select {
		case <-goWatcher.Events:
			compileGo(output, sources)
		case e := <-goWatcher.Errors:
			fmt.Println("ERR", e)
		}
	}
}