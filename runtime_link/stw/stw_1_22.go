//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

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
	reason stwReason
	start  int64
}

//go:linkname stopTheWorld runtime.stopTheWorld
func stopTheWorld(reason stwReason) worldStop

//go:linkname startTheWorld runtime.startTheWorld
func startTheWorld(w worldStop)
