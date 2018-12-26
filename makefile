wasm: main.go exec.go loader.go config.go run.go cssc.go watcher.go goc.go
	go build -o target/wasm $^
