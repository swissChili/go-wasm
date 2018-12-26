package main

import (
	"fmt"
	"io/ioutil"
	"github.com/go-yaml/yaml"
	. "github.com/logrusorgru/aurora"
	// am I really importing something just to copy a dir?
	// yep. 
	. "github.com/otiai10/copy"
	"strings"
	fp "path/filepath"
	"os"
	"log"
	"flag"
	"net/http"
	"github.com/fsnotify/fsnotify"
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

var (
	port = flag.String("port", ":8080", "listen address")
	serve = flag.Bool("serve", true, "start a server?")
	// watches for CSS changes
	watcher   *fsnotify.Watcher
	goWatcher *fsnotify.Watcher
)



func main() {
	flag.Parse()


	dat, err := ioutil.ReadFile("wasm.yml")
	fileNotFound(err)

	var wasm Config
	yaml.Unmarshal(dat, &wasm)

	// Start up both the watchers
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	goWatcher, _ = fsnotify.NewWatcher()
	defer goWatcher.Close()

	// Move all of the static dir contents
	
	status("Copying Static", fmt.Sprintf("%s -> %s", wasm.Static, wasm.Output))
	err = os.RemoveAll(wasm.Output)
	if err != nil { panic(err) }

	err = Copy(wasm.Static, wasm.Output)
	if err != nil { panic(err) }

	sources := strings.Join(wasm.Source, " ")

	for _, s := range wasm.Source {
		if goWatcher.Add(s) != nil {
			panic("Failed to add source "+s+" to watcher")
		}
	}

	// Compile the go files
	compileGo(wasm.Output, sources)

	// Compile all the sass files
	err = compileCSS(wasm.CssDir, wasm.Output, wasm.CssComp)
	if err != nil { panic(err) }

	// Register the dir items to be watched
	err = fp.Walk(wasm.CssDir, watchDir)
	if err != nil { panic(err) }

	// Generate the JavaScript to load 
	const fullLoader = execScript + wasmLoader
	status("Generating JS", wasm.Output + "/index.js")

	err = ioutil.WriteFile(wasm.Output + "/index.js", []byte(fullLoader), 0644)
	if err != nil {
		panic(err)
	}

	// Start server
	
	if *serve {
		done := make(chan bool)

		go watchAndRecompileCSS(wasm.CssDir, wasm.Output, wasm.CssComp)
		go watchAndRecompileGo(wasm.Output, sources)
		status("Serving", fmt.Sprintf("%s on %s", wasm.Output, *port))
		log.Fatal(http.ListenAndServe(*port, http.FileServer(http.Dir(wasm.Output))))

		<-done
	}
}