package app_test

import (
	"testing"
	"unsafe"

	"github.com/rvolosatovs/go-wasm-test/app"
	"github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/http/types"
	"github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/io/streams"
	"github.com/ydnar/wasm-tools-go/cm"
)

var T *testing.T

//go:linkname wasmimport_OutputStreamResourceDrop github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/io/streams.wasmimport_OutputStreamResourceDrop
func wasmimport_OutputStreamResourceDrop(self0 uint32) {}

//go:linkname wasmimport_OutputStreamBlockingWriteAndFlush github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/io/streams.wasmimport_OutputStreamBlockingWriteAndFlush
func wasmimport_OutputStreamBlockingWriteAndFlush(self0 uint32, contents0 *uint8, contents1 uint32, result *cm.Result[streams.StreamError, struct{}, streams.StreamError]) {
	const EXPECTED = "goodbye, world"

	got := unsafe.String(contents0, contents1)
	if got != EXPECTED {
		T.Fatalf("expected: `%s`\ngot: `%s`", EXPECTED, got)
	}
}

//go:linkname wasmimport_NewFields github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/http/types.wasmimport_NewFields
func wasmimport_NewFields() (result0 uint32) {
	return 0
}

//go:linkname wasmimport_ResponseOutparamSet github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/http/types.wasmimport_ResponseOutparamSet
func wasmimport_ResponseOutparamSet(param0 uint32, response0 uint32, response1 uint32, response2 uint32, response3 uint64, response4 uint32, response5 uint32, response6 uint32, response7 uint32) {
}

//go:linkname wasmimport_NewOutgoingResponse github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/http/types.wasmimport_NewOutgoingResponse
func wasmimport_NewOutgoingResponse(headers0 uint32) (result0 uint32) {
	return 0
}

//go:linkname wasmimport_OutgoingResponseBody github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/http/types.wasmimport_OutgoingResponseBody
func wasmimport_OutgoingResponseBody(self0 uint32, result *cm.Result[types.OutgoingBody, types.OutgoingBody, struct{}]) {
}

//go:linkname wasmimport_OutgoingBodyFinish github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/http/types.wasmimport_OutgoingBodyFinish
func wasmimport_OutgoingBodyFinish(this0 uint32, trailers0 uint32, trailers1 uint32, result *cm.Result[types.ErrorCode, struct{}, types.ErrorCode]) {
}

//go:linkname wasmimport_OutgoingBodyWrite github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/http/types.wasmimport_OutgoingBodyWrite
func wasmimport_OutgoingBodyWrite(self0 uint32, result *cm.Result[streams.OutputStream, streams.OutputStream, struct{}]) {
}

func TestHandle(t *testing.T) {
	T = t
	app.Handle(0, 0)
}
