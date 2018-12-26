wasm: main.go exec.go loader.go config.go run.go
	go build -o target/wasm $^
