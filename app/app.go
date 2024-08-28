//go:generate go run github.com/ydnar/wasm-tools-go/cmd/wit-bindgen-go@v0.1.5 generate -w service -o bindings ./wit

package app

import (
	"log/slog"
	"unsafe"

	"github.com/rvolosatovs/go-wasm-test/app/bindings/wasi/http/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

func Handle(request types.IncomingRequest, responseOut types.ResponseOutparam) {
	res := types.NewOutgoingResponse(types.NewFields())

	body := res.Body()
	if body.IsErr() {
		types.ResponseOutparamSet(responseOut, cm.Err[cm.Result[types.ErrorCodeShape, types.OutgoingResponse, types.ErrorCode]](types.ErrorCodeInternalError(cm.None[string]())))
		return
	}
	bodyOut := body.OK()

	bodyWrite := bodyOut.Write()
	if bodyWrite.IsErr() {
		types.ResponseOutparamSet(responseOut, cm.Err[cm.Result[types.ErrorCodeShape, types.OutgoingResponse, types.ErrorCode]](types.ErrorCodeInternalError(cm.None[string]())))
		return
	}

	types.ResponseOutparamSet(responseOut, cm.OK[cm.Result[types.ErrorCodeShape, types.OutgoingResponse, types.ErrorCode]](res))
	stream := bodyWrite.OK()
	s := "hello, world"
	writeRes := stream.BlockingWriteAndFlush(cm.NewList(unsafe.StringData(s), uint(len(s))))
	if writeRes.IsErr() {
		slog.Error("failed to write to stream", "err", writeRes.Err())
		return
	}
	stream.ResourceDrop()

	finishRes := types.OutgoingBodyFinish(*bodyOut, cm.None[types.Fields]())
	if finishRes.IsErr() {
		slog.Error("failed to finish outgoing body", "err", finishRes.Err())
		return
	}
}
