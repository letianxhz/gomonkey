//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stw

import (
	_ "unsafe"
)

func newSTWCtx() ctx {
	return &stwCtx{}
}

type stwCtx struct{}

const stwForTestResetDebugLog = 16

func (ctx *stwCtx) StopTheWorld() {
	stopTheWorld(stwForTestResetDebugLog)
}

func (ctx *stwCtx) StartTheWorld() {
	startTheWorld()
}

//go:linkname stopTheWorld runtime.stopTheWorld
func stopTheWorld(reason uint8)

//go:linkname startTheWorld runtime.startTheWorld
func startTheWorld()
