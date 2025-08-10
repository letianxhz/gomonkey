//go:build !go1.21
// +build !go1.21

package stw

import (
	_ "unsafe"
)

func newSTWCtx() ctx {
	return &stwCtx{}
}

type stwCtx struct{}

func (ctx *stwCtx) StopTheWorld() {
	stopTheWorld("mockey")
}

func (ctx *stwCtx) StartTheWorld() {
	startTheWorld()
}

//go:linkname stopTheWorld runtime.stopTheWorld
func stopTheWorld(reason string)

//go:linkname startTheWorld runtime.startTheWorld
func startTheWorld()
