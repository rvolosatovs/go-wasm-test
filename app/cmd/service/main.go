package main

import (
	"github.com/rvolosatovs/go-wasm-test/app"
	incominghandler "github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/http/incoming-handler"
)

func init() {
	incominghandler.Exports.Handle = app.Handle
}

func main() {}
