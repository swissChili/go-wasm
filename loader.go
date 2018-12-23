package main

const wasmLoader = `
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(res => {
    go.run(res.instance);
});`