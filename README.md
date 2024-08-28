# go-wasm-test

PoC for testing Go Wasm components

Testing:
```
$ go test ./app
```

To build a component:
```
$ cd ./app
$ go generate ./... # optional step, generated code already checked in
$ tinygo build -target=wasip2 -wit-package ./wit -wit-world service ./cmd/service
$ wasmtime serve -Scommon ./service.wasm
```
