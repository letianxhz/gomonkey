//go:build go1.23 && mockey_stw
// +build go1.23,mockey_stw

package stw

import (
	_ "unsafe"
)

func newSTWCtx() ctx {
	return &stwCtx{}
}

type stwCtx struct {
	w worldStop
}

const stwForTestResetDebugLog = 16

func (ctx *stwCtx) StopTheWorld() {
	ctx.w = stopTheWorld(stwForTestResetDebugLog)
}

func (ctx *stwCtx) StartTheWorld() {
	startTheWorld(ctx.w)
}

// stwReason is an enumeration of reasons the world is stopping.
type stwReason uint8

// worldStop provides context from the stop-the-world required by the
// start-the-world.
type worldStop struct {
	reason           stwReason
	startedStopping  int64
	finishedStopping int64
	stoppingCPUTime  int64
}

//go:linkname stopTheWorld runtime.stopTheWorld
func stopTheWorld(reason stwReason) worldStop

//go:linkname startTheWorld runtime.startTheWorld
func startTheWorld(w worldStop)
