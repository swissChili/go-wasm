wasm: main.go exec.go loader.go config.go
	go build -o target/wasm $^
